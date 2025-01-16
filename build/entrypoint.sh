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
if [ ! -z ${POSTGRES_USER} ]; then
  CMD_ARGS="${CMD_ARGS} -db-username ${POSTGRES_USER}"
fi
if [ ! -z ${POSTGRES_PASSWORD} ]; then
  CMD_ARGS="${CMD_ARGS} -db-pass ${POSTGRES_PASSWORD}"
fi
if [ ! -z ${POSTGRES_DB} ]; then
  CMD_ARGS="${CMD_ARGS} -db-name ${POSTGRES_DB}"
fi
if [ ! -z ${DB_SSL} ]; then
  CMD_ARGS="${CMD_ARGS} -db-ssl ${DB_SSL}"
fi
if [ ! -z ${API_URL} ]; then
  CMD_ARGS="${CMD_ARGS} -external-url ${API_URL}"
fi
if [ ! -z ${LOG_LVL} ]; then
  CMD_ARGS="${CMD_ARGS} -log-lvl ${LOG_LVL}"
fi

echo "CMD_ARGS: $CMD_ARGS"

./musicApp $CMD_ARGS &
PID=$!

trap 'finish $PID' TERM INT

wait $PID
