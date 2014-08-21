$(function(){

  // url prefix (dosen't include /)
  var urlPrefix
  // wiki name of this page
  var wikiName

  function init() {
    getSyaroInfo()
    setupUi()
    bindButtons()
  }

  function getSyaroInfo() {
    // get script path
    // http://stackoverflow.com/questions/2161159/get-script-path
    var scripts = document.getElementsByTagName('script')
    var mypath = scripts[scripts.length-1].src

    // get url prefix
    // path = http://hostname/URL_PREFIX/js/dikkinger.js
    var re = /^http:\/\/[^\/]+\/([^\/]+)\/js\//
    urlPrefix = mypath.match(re)[1]

    re = new RegExp(urlPrefix + '(\/.+)$')
    wikiName = location.href.match(re)[1].split('?')[0].split('#')[0]
  }

  function setupUi() {
    $('.alert').hide()
    $('#renameModalInput').val(wikiName)
  }

  function bindButtons() {

    $('#createModalButton').on('click', function() {
      var name = $('#createModalInput').val()
      if (name === "") {
        $('#createErrorAlert').html('<strong>Error</strong> Please fill brank form.')
        $('#createErrorAlert').show()
        return
      }
      if (name[0] !== '/') { name = '/' + name }

      var reqUrl = location.href.replace(wikiName,
          encodeURIComponent(name).replace(/%2F/g, '/'))

      var req = new XMLHttpRequest()
      req.open('GET', reqUrl + '?action=create')

      req.onreadystatechange = function() {
        if (req.readyState === 4) {
          $('#createModalButton').button('reset')

          switch (req.status) {
          case 200:
            // redirect to editor
            location.href = reqUrl + '?view=editor'
            break

          default:
            // show error alert
            $('#createErrorAlert').html('<strong>Error</strong> ' + req.statusText)
            $('#createErrorAlert').show()
            break
          }
        }
      }

      req.send()
      $('#createModalButton').button('loading')
    })

    $('#renameModalButton').on('click', function() {
      var name = $('#renameModalInput').val()
      if (name === "") {
        $('#renameErrorAlert').html('<strong>Error</strong> Please fill brank form.')
        $('#renameErrorAlert').show()
        return
      }
      if (name[0] !== '/') { name = '/' + name }

      var reqUrl = location.href.replace(wikiName,
          encodeURIComponent(name).replace(/%2F/g, '/'))

      var req = new XMLHttpRequest()
      req.open('GET', reqUrl + '?action=rename&oldpath=' + encodeURIComponent(wikiName))

      req.onreadystatechange = function() {
        if (req.readyState === 4) {
          $('#renameModalButton').button('reset')

          switch (req.status) {
          case 200:
            // redirect to page
            location.href = reqUrl
            break

          default:
            // show error alert
            $('#renameErrorAlert').html('<strong>Error</strong> ' + req.statusText)
            $('#renameErrorAlert').show()
            break
          }
        }
      }

      req.send()
      $('#renameModalButton').button('loading')
    })

    $('#deleteModalButton').on('click', function() {
      var reqUrl = location.href

      var req = new XMLHttpRequest()
      req.open('GET', reqUrl + '?action=delete')

      req.onreadystatechange = function() {
        if (req.readyState === 4) {
          $('#deleteModalButton').button('reset')

          switch (req.status) {
          case 200:
            $('#deleteErrorAlert').hide()
            // show success alert
            $('#deleteSuccessAlert').show()
            break

          default:
            // show error alert
            $('#deleteErrorAlert').html('<strong>Error</strong> ' + req.statusText)
            $('#deleteErrorAlert').show()
            break
          }
        }
      }

      req.send()
      $('#deleteModalButton').button('loading')
    })
  }

  init()

})