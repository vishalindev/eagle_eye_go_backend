# Eagle Eye Microservices Endpoint Registry
This repository uses a Go workspace with one module per microservice plus a shared corporate module.

## Services
- `services/admin-service`
- `services/camera-service`
- `services/dashboard-service`
- `services/identity-service`
- `services/mobile-service`
- `services/notification-service`
- `services/platformuser-service`
- `services/reports-service`
- `services/rule-service`
- `services/zone-service`

## Legacy Endpoint Ownership

### admin-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_SAVE_EMPLOYEE` | `POST` | `/Employee/createuser` |
| `API_UPDATE_EMPLOYEE` | `PUT` | `/Employee/updateuser` |
| `API_GET_EMPLOYEES` | `GET` | `/Employee/fetchusers` |
| `API_SAVE_PLANTSETUP` | `POST` | `/Plant/createplant` |
| `API_UPDATE_PLANTSETUP` | `PUT` | `/Plant/updateplant` |
| `API_GET_PLANTSETUPS` | `GET` | `/Plant/fetchplant` |

### platformuser-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_SAVE_PLATFORMUSER` | `POST` | `/PlatFormUser` |
| `API_UPDATE_PLATFORMUSER` | `PUT` | `/PlatFormUser/updateplatformuser` |
| `API_GET_PLATFORMUSERS` | `GET` | `/PlatFormUser/fetchallplatformusers` |
| `API_GET_ROLES` | `GET` | `/PlatFormUser/fetchroleall` |
| `API_GET_DESIGNATIONS` | `GET` | `/PlatFormUser/fetchdesignationall` |

### camera-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_SAVE_CAMERA` | `POST` | `/Camera/addcamera` |
| `API_UPDATE_CAMERA` | `PUT` | `/Camera/updatecamera` |
| `API_GET_CAMERAS` | `GET` | `/Camera/fetchcamerasall` |
| `API_GET_CAMERA_UNITS` | `GET` | `/Camera/departmentwithcamera` |

### zone-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_SAVE_ZONE` | `POST` | `/Zone/createzone` |
| `API_UPDATE_ZONE` | `PUT` | `/Zone/updatezone` |
| `API_GET_ZONES` | `GET` | `/Zone/fetchzoneall` |
| `API_GET_CAMERABYZONEID` | `GET` | `/Zone/fetchzonebyid` |
| `API_GET_ZONE_CODE` | `GET` | `/Zone/createzonecode` |
| `API_UPDATE_ZONE_CAMERA` | `PUT` | `/Zone/zonecameramap` |
| `API_UPDATE_ZONE_RULE` | `PUT` | `/Zone/zonerulemap` |
| `API_GET_ZONE_USER_TAB` | `GET` | `/Zone/fetchzonebytabid` |

### rule-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_GET_RULES` | `GET` | `/Rule/fetchruleall` |
| `API_SAVE_RULE` | `POST` | `/Rule/addnotificationrule` |
| `API_UPDATE_RULE` | `PUT` | `/Rule/updaterule` |
| `API_GET_RULEBYUNITID` | `GET` | `/Rule/fetchrulebyunitid` |

### notification-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_SAVE_NOTIFICATIONGROUP` | `POST` | `/NotificationGroup/addnotificationgroup` |
| `API_UPDATE_NOTIFICATIONGROUP` | `PUT` | `/NotificationGroup/updatenotificationgroup` |
| `API_GET_NOTIFICATIONGROUPS` | `GET` | `/NotificationGroup/fetchnotificationgroupall` |
| `API_GET_NOTIFICATIONSBYCAMERAID` | `GET` | `/Notification/fetchnotificationbycameraid` |
| `API_GET_NOTIFICATIONBYID` | `GET` | `/Notification/fetchnotificationbyid` |
| `API_GET_NOTIFICATIONARRAY` | `GET` | `/Notification/getnotificationarray` |
| `API_GET_TODAYSNOTIFICATIONS` | `GET` | `/Notification/fetchnotification` |
| `API_POST_NOTIFICATION_STATUS` | `POST` | `/Notification/notificationstatusupdate` |

### dashboard-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_GET_ACAMERAS` | `GET` | `/Dashboard/fetchanomalycameras` |
| `API_GET_UNITS` | `GET` | `/Dashboard/fetchunitsall` |
| `API_GET_USERTAB` | `GET` | `/Dashboard/fetchusertab` |
| `API_CREATE_USER_TAB` | `POST` | `/Dashboard/createusertab` |
| `API_FETCH_USER_TAB` | `GET` | `/Dashboard/fetchusertab` |
| `API_DELETE_TAB` | `DELETE` | `/Dashboard/deleteusertab/` |

### reports-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_FETCH_NOTIFICATION_BY_DATE` | `GET` | `/Reports/fetchnotificationbydate` |
| `API_FETCH_NOTIFICATION_BY_VEHICLE` | `GET` | `/Reports/fetchnoticiationvehicle` |
| `API_FETCH_NOTIFICATION_BY_FACE` | `GET` | `/Reports/fetchnoticiationface` |

### identity-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_IDENTITY_HEALTH` | `GET` | `/healthz` |

### mobile-service

| Constant | Method | Path |
| --- | --- | --- |
| `API_MOBILE_HEALTH` | `GET` | `/healthz` |
