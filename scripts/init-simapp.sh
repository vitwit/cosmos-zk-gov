#!/usr/bin/env bash

SIMD_BIN=${SIMD_BIN:=$(which simd 2>/dev/null)}
ALICE_MNEMONIC="all soap kiwi cushion federal skirt tip shock exist tragic verify lunar shine rely torch please view future lizard garbage humble medal leisure mimic"
BOB_MNEMONIC="remain then chuckle hockey protect sausage govern curve hobby aisle clinic decline rotate judge this sail broom debris minute buddy buffalo desk pizza invite"
SAI_MNEMONIC="festival borrow upon ritual remind song execute chase toward fan neck subway canal throw nothing ticket frown leave thank become extend balcony strike fame"
TEJA_MNEMONIC="claim infant gather cereal sentence general cheese float hero dwarf miracle oven tide virus question choice say relax similar rice surround deal smooth rival"
UNKNOWN_MNOMONIC="purpose clutch ill track skate syrup cost among piano elegant close chaos come quit orchard acquire plunge hockey swift tongue salt supreme sting night"

if [ -z "$SIMD_BIN" ]; then echo "SIMD_BIN is not set. Make sure to run make install before"; exit 1; fi
echo "using $SIMD_BIN"
if [ -d "$($SIMD_BIN config home)" ]; then rm -r $($SIMD_BIN config home); fi
$SIMD_BIN config set client chain-id demo
$SIMD_BIN config set client keyring-backend test
$SIMD_BIN config set app api.enable true

echo $ALICE_MNEMONIC | $SIMD_BIN keys add alice --recover
echo $BOB_MNEMONIC | $SIMD_BIN keys add bob --recover
echo $SAI_MNEMONIC | $SIMD_BIN keys add sai --recover
echo $TEJA_MNEMONIC | $SIMD_BIN keys add teja --recover
echo $UNKNOWN_MNOMONIC | $SIMD_BIN keys add unknown --recover

$SIMD_BIN init test --chain-id demo
$SIMD_BIN genesis add-genesis-account alice 5000000000stake --keyring-backend test
$SIMD_BIN genesis add-genesis-account bob 5000000000stake --keyring-backend test
$SIMD_BIN genesis add-genesis-account sai 5000000000stake --keyring-backend test
$SIMD_BIN genesis add-genesis-account teja 5000000000stake --keyring-backend test
$SIMD_BIN genesis add-genesis-account unknown 5000000000stake --keyring-backend test

$SIMD_BIN genesis gentx alice 1000000stake --chain-id demo
$SIMD_BIN genesis collect-gentxs
