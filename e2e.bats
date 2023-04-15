#!/usr/bin/env bats

@test "reject because label key is a palindrome" {
  run kwctl run annotated-policy.wasm -r test_data/pod2.json'

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request rejected
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*false') -ne 0 ]
  [ $(expr "$output" : ".*The label key 'level' is a palindrome.*") -ne 0 ]
}

@test "accept because label key is not a palindrome" {
  run kwctl run annotated-policy.wasm -r test_data/pod.json'
  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}
