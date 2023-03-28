#!/usr/bin/env bash

archived tx wasm store \
/Users/hankbreckenridge/git/cosmos/ARC-H1VE/contracts/artifacts/archive_test-aarch64.wasm \
--chain-id localarchive --from val-0 --gas 2000000 -y

archived tx cda register-contract \
"This is a test description" \
1 \
"breckenridgeh2@gmail.com" \
"google.com/more-info" \
"{\"$schema\": \"https://json-schema.org/draft/2019-09/schema\", \"$id\": \"http://example.com/example.json\", \"type\": \"object\", \"default\": {}, \"title\": \"Root Schema\", \"required\": [ \"ownerships\" ], \"properties\": { \"ownerships\": { \"type\": \"array\", \"default\": [], \"title\": \"The ownerships Schema\", \"items\": { \"type\": \"object\", \"default\": {}, \"title\": \"A Schema\", \"required\": [ \"owner\", \"ownership_proportion\" ], \"properties\": { \"owner\": { \"type\": \"string\", \"default\": \"\", \"title\": \"The owner Schema\", \"examples\": [ \"address\" ] }, \"ownership_proportion\": { \"type\": \"integer\", \"default\": 0, \"title\": \"The ownership_proportion Schema\", \"examples\": [ 1 ] } }, \"examples\": [{ \"owner\": \"address\", \"ownership_proportion\": 1 }] }, \"examples\": [ [{ \"owner\": \"address\", \"ownership_proportion\": 1 }] ] }}, \"examples\": [{ \"ownerships\": [{ \"owner\": \"address\", \"ownership_proportion\": 1 }]}]}" \
"google.com/contract-template" \
"google.com/contract-template-schema" \
1 \
"hank" \
"david" \
--fees 0.001token \
--from val-0 \
--chain-id localarchive -y

archived tx identity register-issuer "Test Issuer" "google.com/more-info" \
--chain-id localarchive --from val-0 -y

archived tx identity issue-certificate \
"archive12smx2wdlyttvyzvzg54y2vnqwq2qjatekl5jhc" \
"RTnUkolHfn&l" \
"google.com/metadata-schema" \
"{
	\"legal_first_name\": \"e468c1334b667b6ada7605814ac45b1f146c6aad638cc2050977e18df15043a6\",
	\"legal_middle_name\": \"517f873af0691df74b0d9972460c2886ae382d0122c7dec2b71a89ee053848f2\",
	\"legal_last_name\": \"8b96aedfa85ae6a9d1a2bb80994452a6445172983a35b4f3ca4fdcd9d6c5cf99\",
	\"state\": \"4b5ac7536c154c5a56b001d8d542fc582b52be6fb42c223e61b785cda6515841\",
	\"postal_code\": \"c349023dbdd3bbcac016ca39a84bb25dfe044d1b235b5b0c63287834e6c3320f\",
	\"city\": \"3f8ca12373e22112d5541105636196cd05efa5d4163646d772b58c119030db24\",
	\"street_address\": \"cf858afa676dd46c2a7de4217ed9764eaf59e0eed3b72128b9a1813a99555de6\"
}" \
--fees 0.001token \
--from val-0 \
--chain-id localarchive -y

archived tx identity accept-identity 0 \
--fees 0.001token \
--from val-0 \
--chain-id localarchive -y

archived tx cda create-cda 0 0 \
"google.com/signing-data" \
'{ "ownerships": [{ "owner": "address", "ownership_proportion": 1 }]}' \
"2023-09-10T15:04:05Z" \
"{}" \
--from val-0 \
--chain-id localarchive -y

archived tx cda approve-cda 0 0 \
'{ "ownerships": [{ "owner": "address", "ownership_proportion": 1 }]}' \
--from val-0 \
--chain-id localarchive \
--fees 0.001token -y
