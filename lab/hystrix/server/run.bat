set GRPC_GO_LOG_SEVERITY_LEVEL=info
rem server.exe -m 1 -c server.toml -p xsf -g xsf -s xsf-server -u "http://10.1.87.69:6868"
server.exe -m 0 -c server.toml -p xsf -g xsf -s xsf-server -u "http://10.1.87.69:6868"