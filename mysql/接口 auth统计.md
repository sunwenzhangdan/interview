<!--
 * @Date: 2021-12-06 16:45:58
 * @LastEditors: sun 
 * @LastEditTime: 2021-12-06 16:45:58
 * @FilePath: 接口 auth统计.md
-->
接口 auth统计
未验证token,参数加密
pcmgr-micro-gateway

POST   /api/v5/micro/new/lock   
POST   /ping                            
POST   /api/v3/superAdmin/pcmgr/upgrade/pcclient/chart/successratio/realtime 
POST   /api/v3/superAdmin/pcmgr/upgrade/pcclient/chart/successratio/28 
POST   /api/v3/superAdmin/pcmgr/upgrade/pcclient/chart/successratio/28/accumulation 
POST   /api/v1/client/id/g     
POST   /api/v1/client/healthweeklies 
GET    /api/v1/client/healthweeklies/date 
POST   /api/v1/client/healthweeklies/data 
POST   /api/v1/client/healthweeklies/research/data 
GET    /api/v1/client/healthweeklies/htmlversion
POST   /api/v1/client/healthweeklies/lservice/data 
POST   /api/v1/client/pcm/ppu   
POST   /api/v1/client/pcm/ppu/r  
POST   /api/v1/client/pcm/recommend/list 
POST   /api/v1/client/pcm/recommend/selected 
POST   /api/v1/client/pcm/ppu/ctr 
POST   /api/v1/client/pcm/sts/mp/v1/g 
POST   /api/v1/client/okr/dr/wlan 
POST   /api/v1/client/pcm/member/baseinfo 
POST   /api/v1/client/pcm/member/growth 
POST   /api/v1/client/message/lewin/status
POST   /api/v1/client/okr/v2/aio 
POST   /api/v1/client/scr/init  
POST   /api/v1/client/pcm/ping  
POST   /api/v1/client/pcm/ping/v6 
POST   /api/v1/client/pcm/sts/sp/v1/l 
POST   /api/v1/client/pcm/sts/sp/init
POST   /api/v1/client/pcm/sts/sp/upgrade/client 
POST   /api/v1/client/pcm/sts/bl 
POST   /api/v1/client/pcm/sts/mktname 
GET    /api/v1/client/pcm/sts/mt/series 
GET    /api/v1/client/pcm/sts/pcmgr/xmlurl 
GET    /api/v1/client/pcm/sts/pcmgr/newuser
POST   /api/v1/client/pcm/sts/sp/v2/init
POST   /api/v1/client/pcm/sts/pcmgr/newuserbyuniqid 


验证token
nodejs-api

POST /api/v2/strategy/loader 
GET  /api/isok 
POST /api/v2/upgrade/client/virusdb 
POST /api/v2/upgrade/client/20200207 
POST /api/v2/messagecenter/message 
POST /api/v2/ProductWarranty/pc 
POST /api/auth 
POST /statistics/upgrade/virus
GET  /api/v2/tools/upgrade
POST /api/v1/log/file
POST /api/lid/v1/login
POST /api/lid/v2/login
POST /api/v1/getAutoAccount
POST /api/v2/dl/rp
POST /api/v2/lenovoservice/specialonLineengineer
GET  /api/v1/client/config/xml
POST /api/v1/statistics/upgrade

ego-gateway

POST /api/m/upgrade/v1/client/report(验证token) 
POST /api/m/lis/v1/wi/state(验证token)
POST /api/m/upgrade/v1/client/headless
POST /api/v2/upgrade/client/20200207
GET  /api/isok
POST /api/m/upgrade/v1/client/ending
POST /api/m/upgrade/v1/client/manul
POST /api/m/upgrade/v1/client/redundance
