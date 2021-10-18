urls="
    42.118.242.143:4000/metrics?host=https://192.168.2.157
    42.118.242.143:4000/metrics?host=https://192.168.2.158
    42.118.242.143:4000/metrics?host=https://192.168.2.164
    42.118.242.143:4000/metrics?host=https://192.168.2.165
    42.118.242.143:4000/metrics?host=https://192.168.2.166
    42.118.242.143:4000/metrics?host=https://192.168.2.167
    42.118.242.143:4000/metrics?host=https://192.168.2.168
    42.118.242.143:4000/metrics?host=https://192.168.2.169
    42.118.242.143:4000/metrics?host=https://192.168.2.170
    42.118.242.143:4000/metrics?host=https://192.168.2.171
    42.118.242.143:4000/metrics?host=https://192.168.2.172
    42.118.242.143:4000/metrics?host=https://192.168.2.173
    42.118.242.143:4000/metrics?host=https://192.168.2.174
    42.118.242.143:4000/metrics?host=https://192.168.2.175
    42.118.242.143:4000/metrics?host=https://192.168.2.176
    42.118.242.143:4000/metrics?host=https://192.168.2.177
    42.118.242.143:4000/metrics?host=https://192.168.2.178
    42.118.242.143:4000/metrics?host=https://192.168.2.179
    42.118.242.143:4000/metrics?host=https://192.168.2.181
    42.118.242.143:4000/metrics?host=https://192.168.2.182
"
for url in $urls; do
   # run the curl job in the background so we can start another job
   # and disable the progress bar (-s)
   echo "fetching $url"
   curl $url &
done
wait

