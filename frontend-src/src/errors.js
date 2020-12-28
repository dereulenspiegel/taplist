export default {
  errorMsg: function(error) {
    var message = ""
    if(error.networkError) {
      if (error.networkError.statusCode === 401) {
        message = 'Unauthorized'
      } else {
        message = "Taplist server down"
      }
    } else if(error.gqlError) {
      message = error.gqlError.message
    } else if(error.graphQLErrors && error.graphQLErrors.lenght > 0) {
      message = error.graphQLErrors[0].message
    } else {
      message = JSON.stringify(error)
    }
    return message
  }
}