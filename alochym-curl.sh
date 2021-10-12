urls="
    localhost:4000/metrics?host=https://192.168.2.157
    localhost:4000/metrics?host=https://192.168.2.158
    localhost:4000/metrics?host=https://192.168.2.164
    localhost:4000/metrics?host=https://192.168.2.165
    localhost:4000/metrics?host=https://192.168.2.166
    localhost:4000/metrics?host=https://192.168.2.167
    localhost:4000/metrics?host=https://192.168.2.168
    localhost:4000/metrics?host=https://192.168.2.169
    localhost:4000/metrics?host=https://192.168.2.170
    localhost:4000/metrics?host=https://192.168.2.171
    localhost:4000/metrics?host=https://192.168.2.172
    localhost:4000/metrics?host=https://192.168.2.173
    localhost:4000/metrics?host=https://192.168.2.174
    localhost:4000/metrics?host=https://192.168.2.175
    localhost:4000/metrics?host=https://192.168.2.176
    localhost:4000/metrics?host=https://192.168.2.177
    localhost:4000/metrics?host=https://192.168.2.178
    localhost:4000/metrics?host=https://192.168.2.179
    localhost:4000/metrics?host=https://192.168.2.181
    localhost:4000/metrics?host=https://192.168.2.182
"
for url in $urls; do
   # run the curl job in the background so we can start another job
   # and disable the progress bar (-s)
   echo "fetching $url"
   curl $url &
done
wait

