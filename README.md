# cicd-practice

Go + [Gin](https://gin-gonic.com/) 애플리케이션과 GitHub Actions / Kind CI 실습 프로젝트입니다.

## 로컬 실행

Go 1.26.5 이상이 필요합니다.

```sh
go mod download
go run .
# 다른 포트로 실행
PORT=8080 go run .
```

- `GET /`: `200`, `Hello World!` (기존 응답 유지)
- `GET /healthz`: `200`, `{"status":"ok"}`
- 기본 포트: `3000` (`PORT` 환경변수로 변경)
- SIGINT/SIGTERM을 받으면 최대 10초 동안 진행 중인 요청을 마무리합니다.

## 테스트 및 빌드

```sh
go test -race -cover ./...
go vet ./...
go build -o bin/server .
./bin/server
```

## Docker

```sh
docker build -t cicd-practice:ci .
docker run --rm -p 3000:3000 cicd-practice:ci
```

멀티 스테이지 빌드로 생성한 Go 바이너리를 비루트 사용자로 실행합니다.

## Kind에서 실행

Docker, Kind, kubectl이 필요합니다.

```sh
kind create cluster --name ci-cluster
docker build -t cicd-practice:ci .
kind load docker-image cicd-practice:ci --name ci-cluster
kubectl --context kind-ci-cluster apply -f k8s/
kubectl --context kind-ci-cluster rollout status deployment/cicd-practice --timeout=120s
kubectl --context kind-ci-cluster port-forward service/cicd-practice 3000:3000
# 다른 터미널
curl http://localhost:3000/
```

매니페스트는 Kind 실습용이며 로컬에 로드한 `cicd-practice:ci` 이미지를 사용합니다 (`imagePullPolicy: Never`). 다른 클러스터에 배포하려면 이미지 레지스트리 주소와 pull policy를 변경하세요.

## CI

`main` 대상 PR과 `main` push마다 Go 포맷, race 테스트, vet, 빌드를 검증합니다. 이후 앱 이미지를 빌드하고 Kind에 로드해 3개 replica의 준비 상태를 확인한 뒤 HTTP 응답을 검증합니다.
