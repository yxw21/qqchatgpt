FROM alpine
WORKDIR /home/qqchatgpt
RUN apk add ca-certificates chromium xvfb && \
    addgroup -S qqchatgpt && adduser -S qqchatgpt -G qqchatgpt -s /bin/ash && \
    version=$(wget -qO- -t1 -T2 "https://api.github.com/repos/yxw21/qqchatgpt/releases/latest" | grep "tag_name" | head -n 1 | awk -F ":" '{print $2}' | sed 's/\"//g;s/,//g;s/ //g') && \
    wget https://github.com/yxw21/qqchatgpt/releases/download/$version/qqchatgpt_linux_amd64.tar.gz && \
    tar -zxvf qqchatgpt_linux_amd64.tar.gz && \
    rm qqchatgpt_linux_amd64.tar.gz README.md && \
    chown -R qqchatgpt:qqchatgpt qqchatgpt && \
    chmod 777 qqchatgpt
USER qqchatgpt
ENTRYPOINT [ "/home/qqchatgpt/qqchatgpt" ]