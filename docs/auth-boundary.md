# 인증 경계

이 문서는 현재 WalkQuest Service의 ExplorerProfile 구현에서 사용하는 인증 경계와 사용자 식별 방식을 정리한다.

## 현재 책임 분리

WalkQuest Service는 JWT를 직접 검증하지 않는다. 인증과 JWT 검증은 Identity/Gateway가 담당하는 것을 전제로 한다.

Gateway는 인증된 사용자 ID를 `X-User-Id` 헤더로 주입하고, WalkQuest Service는 이 `X-User-Id` 값을 기준으로 사용자를 식별한다.

현재 구현에서 WalkQuest Service는 다음 동작을 수행한다.

- `X-User-Id` 헤더가 없으면 `401 Unauthorized`를 반환한다.
- `X-User-Id` 헤더 값이 공백이면 `401 Unauthorized`를 반환한다.
- `X-User-Id`의 존재 여부와 공백 여부는 확인한다.
- JWT 서명 검증은 중복 수행하지 않는다.

## 신뢰 모델

이 신뢰 모델은 외부 클라이언트가 WalkQuest Service에 직접 접근할 수 없다는 네트워크 경계를 전제로 한다. 운영 환경에서는 클라이언트 요청이 Gateway를 거쳐 WalkQuest Service로 전달되어야 한다.

Gateway는 클라이언트가 보낸 기존 `X-User-Id` 헤더를 제거하고, 인증 결과로 확인된 사용자 ID 값을 새 `X-User-Id` 헤더로 주입해야 한다. WalkQuest Service는 Gateway가 주입한 이 값을 신뢰한다.

따라서 `X-User-Id`를 직접 전달하는 방식은 운영 환경의 최종 인증 방식이 아니다. 이 방식은 로컬 개발과 테스트에서 Gateway를 대신하기 위한 검증 방법이다.

## 현재 구현 상태

현재 Gateway와 Identity 연동은 아직 구현되지 않았다.

현재 로컬 검증에서는 Gateway 대신 요청에 `X-User-Id` 헤더를 직접 전달한다. 이 방식은 로컬 개발과 테스트 용도로만 사용한다.
