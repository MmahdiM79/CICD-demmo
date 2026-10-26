# cicd-demo

A tiny app for live CI/CD demos. Written in Go with only the standard library; no Node.js or npm required.

```
push to main ──► 🧪 test ──► 🐳 build & push image (ghcr.io) ──► 🚀 deploy (GitHub Pages)
pull request ──► 🧪 test only
```

## Prerequisites

- Go 1.27 or newer (`go version`).
- Git and a GitHub account.
- Docker Desktop or OrbStack for Demo 2.
- Optional: GitHub CLI (`gh`); the Actions tab works too.
- Optional rolling update: kind and kubectl.
- For an offline Docker demo, pull `golang:1.27-alpine` and `alpine:3.23` and build the image before class.

Local server: `go run ./cmd/server` (http://localhost:3000).
Static export: `go run ./cmd/static` (writes `dist/index.html`).

## One-time setup (do this before class)

1. Create a **public** GitHub repo called `cicd-demo` and push this folder to it:
   ```bash
   cd demo
   git init -b main && git add . && git commit -m "init: cicd demo"
   git remote add origin git@github.com:<you>/cicd-demo.git
   git push -u origin main
   ```
2. **Settings → Pages → Build and deployment → Source: GitHub Actions**.
3. Re-run the workflow (Actions → CI/CD Pipeline → Run workflow). The site goes live at
   `https://<you>.github.io/cicd-demo/`.
4. Optional, to show *Continuous Delivery*: **Settings → Environments → github-pages →
   Required reviewers → add yourself**. Every deploy then waits for your approval.
5. Optional, to show a red check blocking a merge: **Settings → Branches → add rule for `main` →
   Require status checks → `🧪 Test`**.

## Demo 1: your first pipeline
```bash
go test -v ./...
# edit Headline in internal/demo/config.go
git commit -am "feat: hello class" && git push
gh run watch                              # or open the Actions tab
```

## Demo 2: build, ship, go live
```bash
docker build -t demo --build-arg APP_VERSION=local .
docker run --rm -p 3000:3000 demo         # http://localhost:3000
docker build -t demo .                    # run again → layers are CACHED
# change Accent / Headline in internal/demo/config.go
git commit -am "feat: new look" && git push
```
Then show the image under **Packages** (tagged with the commit SHA) and open the Pages URL.

## Demo 3: break the build
```bash
# internal/demo/pricing.go:  (1-discount)  →  (1+discount)
git commit -am "tweak discount" && git push    # 🧪 test fails, build + deploy are skipped
git revert HEAD --no-edit && git push          # green again → that's your MTTR
```

## Optional: rolling update on a local Kubernetes cluster
```bash
kind create cluster
docker build -t cicd-demo:v1 --build-arg APP_VERSION=v1 . && kind load docker-image cicd-demo:v1
kubectl apply -f k8s/
docker build -t cicd-demo:v2 --build-arg APP_VERSION=v2 . && kind load docker-image cicd-demo:v2
kubectl set image deploy/cicd-demo app=cicd-demo:v2
kubectl get pods -w                      # pods are replaced one by one
kubectl rollout undo deploy/cicd-demo    # rollback
```

## Files
| Path | What |
|---|---|
| `go.mod` | Go version; no third-party modules |
| `internal/demo/config.go` | headline and accent colour |
| `internal/demo/pricing.go` | discount logic for Demo 3 |
| `internal/demo/page.go` | shared HTML rendering |
| `internal/demo/*_test.go` | pricing, page and HTTP tests |
| `cmd/server/main.go` | HTTP server with graceful shutdown |
| `cmd/static/main.go` | exports the same page to GitHub Pages |
| `Dockerfile` | multi-stage build; final image contains the Go binary |
| `.github/workflows/ci-cd.yml` | test, build image and deploy |
| `k8s/deployment.yaml` | optional rolling-update demo |
