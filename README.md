# walkquest-service

WalkQuest는 사용자의 실제 이동 기록을 기반으로 생활권 주변 지역을 탐험하고, 방문한 지역을 기록으로 쌓아가는 위치 기반 탐험 서비스입니다.

이 저장소는 WalkQuest 서비스의 Go 기반 백엔드 구현을 위한 프로젝트입니다.

## 프로젝트 목적

이 프로젝트는 Go 기반 백엔드 포트폴리오에서 실제 서비스 도메인 역할을 담당합니다.

Identity/Gateway/Async Platform과 연동 가능한 독립 서비스로 설계하며 `사용자별 데이터 관리`, `탐험 세션`, `GPS 포인트 저장`, `방문 지역 기록` 흐름을 구현하는 것을 목표로 합니다.

## MVP 목표

MVP에서는 캐릭터 성장, 보상, 소셜 기능보다 탐험 세션과 GPS 포인트 저장 이후에 방문 지역 판정을 위한 백엔드 기반을 우선 구현할 예정입니다.

## MVP 예정 범위

MVP에서 다룰 예정인 범위는 다음과 같습니다.

- 탐험 프로필 생성 및 조회
- 탐험 세션 시작 및 조회
- 진행 중 세션 제한
- GPS 포인트 업로드
- 세션 상세 조회
- 좌표 기반 방문 지역 판정
- 방문 기록 저장
- 발견 로그 생성
- 세션 종료 멱등성

## MVP 제외 범위

이번 MVP에서는 다음 기능을 제외합니다.

- 캐릭터 성장
- 경험치 및 보상
- 펫
- 수집
- 친구 및 소셜
- 채팅
- 랭킹
- 장소 추천
- 맛집 추천
- 쿠폰
- Android 앱
- 복잡한 지도 UI

## 기술 스택

- Language: Go
- HTTP: 현재 `net/http` 사용, 추후 Gin 전환 예정
- Database: PostgreSQL
- Infra: Docker Compose
- Auth Boundary: Gateway-provided `X-User-Id` 연동 예정
- Async Integration: 추후 Async Platform과 이벤트 연동 예정

## 현재 구현 상태

현재 구현된 내용은 다음과 같습니다.

- 프로젝트 기본 구조 생성
- `/health` API 구현
- `PORT` 환경변수 기반 서버 실행 설정
- PostgreSQL Docker Compose 구성
- 서버 시작 시 DB ping 확인
- 로컬 실행 및 health check 확인

## 실행 방법

PostgreSQL을 실행합니다.

```bash
docker-compose up -d
```

서버를 실행합니다.

```bash
go run ./cmd/server
```

서버는 기본적으로 `8080` 포트에서 실행됩니다.

서버 시작 시 PostgreSQL ping을 확인하고, 성공하면 로그를 남깁니다.

다른 포트를 사용하려면 `PORT` 환경변수를 설정합니다. 이후 아래와 같이 실행합니다.

```bash
PORT=8081 go run ./cmd/server
```

health check를 확인합니다.

```bash
curl http://localhost:8080/health
```

Windows PowerShell에서는 다음 명령어를 사용할 수 있습니다.

```powershell
curl.exe http://localhost:8080/health
```

예상 응답은 다음과 같습니다.

```json
{
  "status": "ok",
  "service": "walkquest-service"
}
```

## 환경변수 설정

애플리케이션은 실행 시 OS 환경변수를 읽고, 값이 없으면 코드에 정의된 기본값을 사용합니다.

현재 `.env` 파일은 자동으로 로드하지 않습니다.  
`.env.example`은 자동 로딩되는 설정 파일이 아니라, 사용할 수 있는 환경변수 목록을 보여주는 예시 파일입니다.

기본 DB 설정은 로컬 개발용 Docker Compose 설정과 일치하도록 구성되어 있습니다.

## 관련 문서

추후 다음 문서를 추가할 예정입니다.

- `docs/domain-model.md`: 핵심 도메인 모델
- `docs/api.md`: API 명세
- `docs/adr/`: 주요 설계 결정
- `docs/policies.md`: 세션, GPS, 탐험 판정 정책
- `docs/auth-boundary.md`: 인증 경계와 `X-User-Id` 신뢰 모델

## ExplorerProfile 로컬 실행 및 검증

이 섹션은 현재 ExplorerProfile 2주차 구현 상태를 로컬에서 실행하고 검증하는 방법을 정리한다.

### 1. PostgreSQL 실행

```bash
docker compose up -d
```

### 2. 컨테이너 상태 확인

```bash
docker compose ps
```

### 3. migration 적용

```bash
docker compose exec -T postgres psql -U walkquest -d walkquest < migrations/001_create_explorer_profiles.sql
```

이 명령은 현재 실행 중인 `postgres` 컨테이너에 migration SQL 파일을 전달해 실행한다. 이미 같은 migration을 적용한 상태에서 다시 실행하면 테이블이 이미 존재한다는 오류가 발생할 수 있다.

현재는 별도의 migration 관리 도구 없이 SQL 파일을 수동 적용한다.

### 4. DB 콘솔 접속

```bash
docker compose exec postgres psql -U walkquest -d walkquest
```

대표적인 확인 명령은 다음과 같다.

```sql
\dt
\d explorer_profiles
SELECT * FROM explorer_profiles;
\q
```

### 5. 서버 실행

기본 실행 명령은 다음과 같다.

```bash
go run ./cmd/server
```

기본 포트는 `8080`이다.

기본 포트를 사용할 수 없는 경우에만 환경변수로 임시 포트를 지정한다.

PowerShell:

```powershell
$env:PORT="18080"
go run ./cmd/server
```

Git Bash:

```bash
PORT=18080 go run ./cmd/server
```

임시 포트를 사용했다면 아래 API URL의 `8080` 부분도 같은 포트로 바꿔야 한다. 예를 들어 `PORT=18080`을 사용했다면 API URL은 `http://localhost:18080`을 사용한다.

### 6. PowerShell의 curl 사용 주의

Windows PowerShell에서 `curl`은 환경에 따라 `Invoke-WebRequest` 별칭으로 동작할 수 있다. README의 검증 명령은 `curl.exe` 기준으로 작성한다.

Git Bash나 일반 터미널에서는 환경에 따라 `curl`을 사용할 수 있다.

### 7. API 수동 검증

아래 예시는 기본 포트인 `8080` 기준이다.

#### Health check

```powershell
curl.exe -i http://localhost:8080/health
```

기대 결과:

- `200 OK`

#### X-User-Id 없이 프로필 생성

```powershell
curl.exe -i -X POST http://localhost:8080/explorer-profiles
```

기대 결과:

- `401 Unauthorized`
- `{"error":"unauthorized"}`

#### 프로필 생성 성공

```powershell
curl.exe -i -X POST http://localhost:8080/explorer-profiles -H "X-User-Id: user-123"
```

기대 결과:

- `201 Created`
- 생성된 프로필 JSON 반환

#### 내 프로필 조회 성공

```powershell
curl.exe -i http://localhost:8080/explorer-profiles/me -H "X-User-Id: user-123"
```

기대 결과:

- `200 OK`
- 프로필 JSON 반환

#### 동일 사용자 프로필 중복 생성

```powershell
curl.exe -i -X POST http://localhost:8080/explorer-profiles -H "X-User-Id: user-123"
```

기대 결과:

- `409 Conflict`
- `{"error":"explorer_profile_already_exists"}`

#### 존재하지 않는 사용자 프로필 조회

```powershell
curl.exe -i http://localhost:8080/explorer-profiles/me -H "X-User-Id: user-not-exists"
```

기대 결과:

- `404 Not Found`
- `{"error":"explorer_profile_not_found"}`

### 8. 자동 테스트 실행

```bash
go test ./...
```

Service 테스트는 fake repository를 사용해 repository 함수 호출, `userID` 전달, 프로필 반환, 도메인 에러 전달을 검증한다.

Handler 테스트는 fake service와 `httptest`를 사용해 HTTP 상태 코드, JSON 응답, 라우팅, service 호출을 검증한다.

현재 Repository의 실제 PostgreSQL 동작은 수동 통합 검증으로 확인했다. 이번 2주차 범위에서는 별도의 Repository 자동 통합 테스트나 test container는 추가하지 않았다.
