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
