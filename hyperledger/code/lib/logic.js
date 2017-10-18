/**
 * Set User wallet function.
 */

function onSetWallet(setWallet) {
    setWallet.user2.wallet = setWallet.newWallet; 
    return getAssetRegistry('world.alphabets.User')
      .then(function (assetRegistry) {
          return assetRegistry.update(setWallet.user2);
      });
}

/**
 * Set Bet status function.
 */

function onSetPaid(setPaid) {
    setPaid.valuePaid.paid = setPaid.newPaidValue; 
    return getAssetRegistry('world.alphabets.Bet')
      .then(function (assetRegistry) {
          return assetRegistry.update(setPaid.valuePaid);
      });
}