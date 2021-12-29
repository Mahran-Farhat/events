package models

import (
    "time"

)

type Screen struct {
    EventType string `json:"eventType"`
    ScreenName string `json:"screenName"`
    ScreenStartTime time.Time `json:"screenStartTime"`
    ScreenEndTime time.Time `json:"screenEndTime"`
    TotalScreenTime time.Time `json:"totalScreenTime"`
}

type Page struct {
    EventType string `json:"eventType"`
    PageName string `json:"pageName"`
    PageStartTime time.Time `json:"pageStartTime"`
    PageEndTime time.Time `json:"pageEndTime"`
    TotalPageTime time.Time `json:"totalPageTime"`
}

type Section struct {
    EventType string `json:"eventType"`
    SectionName string `json:"sectionName"`
    SectionStartTime time.Time `json:"sectionStartTime"`
    SectionEndTime time.Time `json:"sectionEndTime"`
    TotalSectionTime time.Time `json:"totalSectionTime"`
}

type DeviceInformation struct {
  AppInstalledTime int64  `json:"appInstalledTime"`
  AppVersion string `json:"appVersion"`
  DeviceModel string `json:"deviceModel"`
  Manufacturer string `json:"manufacturer"`
  SystemName string `json:"systemName"`
  SystemVersion string `json:"systemVersion"`
}

type NavigationParams struct{
 Id string `json:"id"`
 Typeparam string `json:"type"`
}

type ScreenEvent struct{
 DeviceInformation DeviceInformation `json:"deviceInformation"`
 EventType string `json:"eventType"`
 NavigationParams NavigationParams `json:"navigationParams"`
 ScreenEndTime int64 `json:"screenEndTime"`
 ScreenName string `json:"screenName"`
 ScreenStartTime int64 `json:"screenStartTime"`
 SessionId string `json:"sessionId"`
 TotalScreenTime int64 `json:"totalScreenTime"`
 UserId string `json:"userId"`
}

type PageEvent struct{
 DeviceInformation DeviceInformation `json:"deviceInformation"`
 EventType string `json:"eventType"`
 NavigationParams NavigationParams `json:"navigationParams"`
 PageEndTime int64 `json:"pageEndTime"`
 PageName string `json:"pageName"`
 PageStartTime int64 `json:"pageStartTime"`
 SessionId string `json:"sessionId"`
 TotalPageTime int64 `json:"totalPageTime"`
 UserId string `json:"userId"`
}

type SectionEvent struct{
 DeviceInformation DeviceInformation `json:"deviceInformation"`
 EventType string `json:"eventType"`
 NavigationParams NavigationParams `json:"navigationParams"`
 SectionEndTime int64 `json:"pageEndTime"`
 SectionName string `json:"pageName"`
 SectionStartTime int64 `json:"pageStartTime"`
 SessionId string `json:"sessionId"`
 TotalSectionTime int64 `json:"totalPageTime"`
 UserId string `json:"userId"`
}




