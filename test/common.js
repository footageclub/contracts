const bip39 = require('bip39');
const wallet = require('ethereumjs-wallet');

Object.defineProperty(exports, "__esModule", { value: true });
exports.makePrivateKey = async function(mnemonic, path) {
    const seed = await bip39.mnemonicToSeed(mnemonic); // mnemonic is the string containing the words
    const hdk = wallet.hdkey.fromMasterSeed(seed);
    const addr_node = hdk.derivePath(path); //m/44'/60'/0'/0/0 is derivation path for the first account. m/44'/60'/0'/0/1 is the derivation path for the second account and so on
    const addr = addr_node.getWallet().getAddressString(); //check that this is the same with the address that ganache list for the first account to make sure the derivation is correct
    const key = addr_node.getWallet().getPrivateKey();
    return {addr:addr, key:key.toString('hex')}
}