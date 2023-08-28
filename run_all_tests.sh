# Run all the tests locally except those that hang (pnpm test)
LOG_FILE_NAME=tests.log
echo > $LOG_FILE_NAME
echo "Log file: $LOG_FILE_NAME"
echo "Starting from a fresh state"
make nuke >> $LOG_FILE_NAME 2>&1
echo "Build..."
make >> $LOG_FILE_NAME 2>&1
make build >> $LOG_FILE_NAME 2>&1
echo "Running op-node tests..."
make -C ./op-node test >> $LOG_FILE_NAME 2>&1
echo "Running op-proposer tests..."
make -C ./op-proposer test >> $LOG_FILE_NAME 2>&1
echo "Running op-batcher tests..."
make -C ./op-batcher test >> $LOG_FILE_NAME 2>&1
echo "Running op-e2e tests..."
make -C ./op-e2e test >> $LOG_FILE_NAME 2>&1
echo "All done, please check $LOG_FILE_NAME."


