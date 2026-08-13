# Gin CI/CD Practice

Gin 기반의 간단한 HTTP 애플리케이션과 EKS 배포 구성을 포함합니다.

## 로컬 실행

```bash
go run .
```

기본 포트는 `3000`이며 `PORT` 환경변수로 변경할 수 있습니다.

```bash
PORT=8080 go run .
```

`GET /` 요청은 `Hello World!`를 반환합니다.

## 테스트

```bash
go test ./...
```

## Docker

```bash
docker build -t gin-app .
docker run --rm -p 3000:3000 gin-app
```

## 배포

`main` 브랜치의 pull request에서는 Docker 이미지의 kind 배포와 HTTP 검증을 수행합니다.
`main` 브랜치에 push하면 검증 후 이미지를 ECR의 `gin-app` 저장소에 올리고 EKS에 배포합니다.
