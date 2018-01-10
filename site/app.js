
$(document).ready(function () {
  //calcualte balance
  var ethAdd = $("#faucet-address").html()
  getBalance(ethAdd)

  $("#request").on("click", function () {
    var value = $("#user_address").val()
    if (!isAddress(value)) {
      $("#error").html("Invalid address. Address should start with 0x, followed by exactly 40 hex digits")
      $("#error").show()
      $("#success").hide();
      return
    }
    submitAddress(value)
  })
})

function getBalance(address){
  $.ajax({
    url: "https://ropsten.etherscan.io/api?module=account&action=balance&address="+address+"&tag=latest",
    type:"GET",
    success: function (response) {
      var balance = response.result
      var wei = Math.pow(10,18)
      var balanceEth = balance/wei
     $("#balance").html(balanceEth)
    }
  })
}

function submitAddress(address) {
  console.log(address)
  $.ajax({
    url: "https://faucet-backend.kyber.network/claim-eth",
    type:"POST",
    data: { address: address },
    success: function (response) {
     // var response = JSON.parse(result)
      if(response.success){
        $("#error").hide()
        $("#success").html(response.msg)
        $("#success").show();
      }else{
        if(response.tx){
          var viewTx = $("<a target='_blank' class='link'>View on etherscan</a>")
          viewTx.attr("href","https://ropsten.etherscan.io/tx/"+response.tx)
          $("#error").html(response.error)
          $("#error").append($("<br>"))
          $("#error").append(viewTx)
        }else{
          var error = response.error?response.error:"Cannot request faucet, Please try at another momment!"
          $("#error").html(error)
        }
        $("#error").show()
        $("#success").hide();
      }
    }
  })
}

/**
 * Checks if the given string is an address
 *
 * @method isAddress
 * @param {String} address the given HEX adress
 * @return {Boolean}
*/
var isAddress = function (address) {
  if (!/^(0x)?[0-9a-f]{40}$/i.test(address)) {
    // check if it has the basic requirements of an address
    return false;
  } else if (/^(0x)?[0-9a-f]{40}$/.test(address) || /^(0x)?[0-9A-F]{40}$/.test(address)) {
    // If it's all small caps or all all caps, return true
    return true;
  } else {
    // Otherwise check each case
    return true
    return isChecksumAddress(address);
  }
};


/**
 * Checks if the given string is a checksummed address
 *
 * @method isChecksumAddress
 * @param {String} address the given HEX adress
 * @return {Boolean}
*/
var isChecksumAddress = function (address) {
  // Check each case
  address = address.replace('0x','');
  var addressHash = sha3(address.toLowerCase());
  for (var i = 0; i < 40; i++ ) {
      // the nth letter should be uppercase if the nth digit of casemap is 1
      if ((parseInt(addressHash[i], 16) > 7 && address[i].toUpperCase() !== address[i]) || (parseInt(addressHash[i], 16) <= 7 && address[i].toLowerCase() !== address[i])) {
          return false;
      }
  }
  return true;
};
