FROM grafana/grafana-oss:9.3.6

##################################################################
## CONFIGURATION
##################################################################

## Set Grafana options
ENV GF_ENABLE_GF_INSTALL_PLUGINS=grafana-opensearch-datasource
ENV GF_USERS_DEFAULT_THEME=light

##################################################################
## VISUAL
##################################################################

## Replace Logo
COPY customizied_login_page/Logo_Celanese.svg /usr/share/grafana/public/img/grafana_icon.svg

## Update Background
COPY customizied_login_page/light-background.svg /usr/share/grafana/public/img/g8_login_dark.svg
COPY customizied_login_page/light-background.svg /usr/share/grafana/public/img/g8_login_light.svg

## Update Title
## Update Login Title
COPY customizied_login_page/build/2362.5e93872490cc5d80351e.js /usr/share/grafana/public/build/ 
COPY customizied_login_page/build/2362.5e93872490cc5d80351e.js.map /usr/share/grafana/public/build/
