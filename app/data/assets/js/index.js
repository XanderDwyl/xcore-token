web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
asciiToHex = Web3.utils.asciiToHex;

contractInstance = new web3.eth.Contract(ABI_DEFINITION, CONTRACT_ADDRESS);

function post() {
  console.log("POSTING");
}

// function post() {
//   var postTitle = $('#postTitle').val();
//   var postBody = $('#postBody').val();

//   var hashedData = "sha512: " + sha512(postTitle + postBody);
//   var timestamp = Math.floor(Date.now() / 1000);

//   var sendData = {
//     "hash": hashedData,
//     "timestamp": timestamp
//   };

//   console.log(JSON.stringify(sendData));

//   $.ajax({
//     type: 'POST',
//     url: "http://localhost:3000/add_comment",
//     data: {
//       "title": postTitle,
//       "body": postBody
//     }
//   }).done(function (res) {
//     web3.eth.getAccounts()
//       .then((accounts) => {
//         console.log("SUCCESS")
//         console.log(accounts)
//         console.log(sendData)
//         return contractInstance.methods.set(JSON.stringify(sendData)).send({
//           from: accounts[0]
//         })
//       })
//       .then(() => {
//         console.log("GET")
//         console.log("SUCCESS")
//         return contractInstance.methods.get().call();
//       })
//       .then((result) => {

//         $('#jsonResult').html(result);
//         console.log(result);

//         $('#postTitle').val("");
//         $('#postBody').val("");
//       });

//   });


  // web3.eth.getAccounts()
  //   .then((accounts) => {
  //     console.log("SUCCESS")
  //     console.log(accounts)
  //     console.log(sendData)
  //     return contractInstance.methods.set(JSON.stringify(sendData)).send({
  //       from: accounts[0]
  //     })
  //   })
  //   .then(() => {
  //     console.log("GET")
  //     console.log("SUCCESS")
  //     return contractInstance.methods.get().call();
  //   })
  //   .then((result) => {

  //     $('#jsonResult').html(result);
  //     console.log(result);

  //     $('#postTitle').val("");
  //     $('#postBody').val("");
  //   });
// }

$(document).ready(function () {
  contractInstance.methods.get().call()
    .then((val) => {
      $('#jsonResult').html(val);
      console.log(val)
    });
});