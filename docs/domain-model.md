# WalkQuest Domain Model

## 1. 목적

이 문서는 WalkQuest MVP에서 사용하는 핵심 도메인 및 각 도메인의 책임을 정리하기 위한 문서입니다.
MVP 개발 과정에서 도메인 책임이 섞이거나 구현 범위가 과도하게 확장되는 것을 방지하고 탐험 중심 MVP 범위를 명확히 유지하는 것을 목표로 합니다.

## 2. 핵심 도메인

- ExplorerProfile
- ActivitySession
- RoutePoint
- ExplorationArea
- VisitedArea
- DiscoveryLog

## 3. 도메인별 책임

### ExplorerProfile

WalkQuest를 사용하는 사람들의 탐험 프로필입니다.
Identity 서비스에서 받는 User 자체를 대체하지 않고 Gateway로부터 받은 `X-User-Id`를 기준으로 WalkQuest 내부 프로필 정보와 연결됩니다.

### ActivitySession

사용자의 한 번의 탐험 진행 세션입니다.
사용자는 탐험을 시작하면 하나의 ActivitySession 안에서 이동 기록을 남기고 이후 세션에 쌓인 RoutePoint를 기반으로 방문 지역을 계산할 수 있습니다.

### RoutePoint

ActivitySession 중 업로드되는 GPS 좌표 데이터입니다.
사용자가 이동하는 동안 기록되는 위도, 경도, 정확도, 기록 시간 등의 위치 정보이며, 이후 세션 종료 시 방문 지역을 계산하기 위한 입력값으로 사용됩니다.

### ExplorationArea

GPS 좌표를 탐험 판정에 사용할 수 있도록 변환한 지도상의 지역 단위입니다.
위도와 경도를 그대로 비교하지 않고 일정한 기준의 지역 단위로 묶어 사용자가 어떤 지역을 탐험했는지 판단하기 위한 기준으로 사용됩니다.

### VisitedArea

사용자가 방문한 ExplorationArea를 기록하는 도메인입니다.
ExplorationArea는 지도상의 지역 단위이며, VisitedArea는 특정 사용자가 해당 지역을 방문했다는 사용자별 방문 기록입니다.

### DiscoveryLog

사용자가 새로운 지역을 처음 발견했을 때 남기는 기록입니다.
VisitedArea 중에서도 신규 방문으로 판단된 경우 생성되며, 사용자가 새로운 지역을 탐험했다는 이벤트성 기록으로 사용됩니다.

## 4. 도메인 관계 개요

ExplorerProfile는 여러 ActivitySession을 가질 수 있습니다.
ActivitySession는 여러 RoutePoint를 가질 수 있습니다.
RoutePoint는 세션 종료 시 ExplorationArea를 계산하기 위한 입력값으로 사용됩니다.
계산된 ExplorationArea 중 사용자가 방문한 지역은 VisitedArea로 기록됩니다.
DiscoveryLog는 사용자가 기존에 방문하지 않았던 신규 지역을 발견했을 때 생성됩니다.

## 5. MVP 제외 도메인

이번 MVP에서는 다음 도메인을 구현하지 않습니다.

- Character
- RewardHistory
- Pet
- Collection
- Social
- Ranking
- PlaceRecommendation

위 항목은 MVP 이후 확장 후보로 분리합니다.

## 6. 다음 설계 주제

- ExplorerProfile 생성/조회 정책
- `X-User-Id` 누락 시 에러 처리
- 동일 사용자의 중복 프로필 생성 정책
