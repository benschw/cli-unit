#!/bin/bash

COUNTER=0
PASSED=0

getBlock() {
	local BODY=$1
	local TYPE="$2"
	local capture="0"

	#for p in $BODY; do
	IFS=''
	while read -r p; do

		if [[ $capture == "1" ]]; then	
			if [[ $p == "###"* ]]; then
				capture="0"
			else
				if [[ "$p" == $'\t'* ]]; then
					trimmed=$(echo "$p" | cut -c 2-)
					echo "$trimmed"
				fi

			fi
		fi
		if [[ $p == "$TYPE"* ]]; then
			capture="1"
		fi

	done <<< "$BODY"
	#done
}

getShell() {
	local BODY="$1"
	getBlock "$BODY" "#### when:"
}
getOutput() {
	local BODY="$1"
	getBlock "$BODY" "#### then:"
}


runTest() {
	local TITLE="$1"
	local BODY="$2"
	COUNTER=$((COUNTER + 1))

	local SHELL=$(getShell "$BODY")
	local EXPECTED=$(getOutput "$BODY")
	local FOUND=$(eval "$SHELL")

	if ! diff <(echo "$EXPECTED") <(echo "$FOUND") > /dev/null; then
		echo "--- FAIL: $TITLE"
		diff <(echo "$EXPECTED") <(echo "$FOUND")
		return 1
	fi
	PASSED=$((PASSED + 1))
	return 0
}

runTests() {
	local FILE=$1

	local title=""
	local unit=""
	local FAILURES=0
	local capture="0"

	IFS=''
	while read -r p; do
		if [[ $capture == "1" ]]; then
			if [[ $p == "### test:"* ]]; then
				capture="0"
				runTest "$title" "$unit" || FAILURES=$((FAILURES + 1))
				unit=""
			else
				unit="$unit$p"$'\n'
			fi
		fi

		if [[ $p == "### test:"* ]]; then
			title=$(echo "$p" | cut -c 10-)
			capture="1"
		fi
	done <$FILE

	runTest "$title" "$unit" || FAILURES=$((FAILURES + 1))
	return $FAILURES
}

FAILURES=0
for f in $(ls $1); do
	runTests $f || FAILURES=$((FAILURES + 1))
done

if [[ $FAILURES > 0 ]]; then
	echo "FAIL ($PASSED/$COUNTER tests passed)"
	exit 1
fi

echo "OK ($PASSED/$COUNTER tests passed)"
