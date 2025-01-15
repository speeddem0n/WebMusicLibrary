#!/bin/sh

finish() {
  ppid=$1
  if [ -n "$ppid" ]; then
    kill -s TERM $ppid
  fi

  # wait until $ppid finished normally to provide more time for termination process
  while [ true ]; do
    ps -ef | grep $ppid | grep -v grep > /dev/null 2>&1
    if [ $? -eq 1 ]; then
      echo "process finished"
      break
    fi
    sleep 2
  done

  return 0
}

CMD_ARGS=
if [ ! -z ${DB_HOST} ]; then
  CMD_ARGS="${CMD_ARGS} -db-host ${DB_HOST}"
fi
if [ ! -z ${DB_PORT} ]; then
  CMD_ARGS="${CMD_ARGS} -db-port ${DB_PORT}"
fi
if [ ! -z ${DB_USERNAME} ]; then
  CMD_ARGS="${CMD_ARGS} -db-username ${DB_USERNAME}"
fi
if [ ! -z ${DB_PASSWORD} ]; then
  CMD_ARGS="${CMD_ARGS} -db-password ${DB_PASSWORD}"
fi
if [ ! -z ${DB_SSL} ]; then
  CMD_ARGS="${CMD_ARGS} -db-ssl ${DB_SSL}"
fi
if [ ! -z ${API_URL} ]; then
  CMD_ARGS="${CMD_ARGS} -api_url ${API_URL}"
fi
if [ ! -z ${LOG_LVL} ]; then
  CMD_ARGS="${CMD_ARGS} -log-lvl ${LOG_LVL}"
fi

echo "CMD_ARGS: $CMD_ARGS"

./musicApp $CMD_ARGS &
PID=$!

trap 'finish $PID' TERM INT

wait $PID
