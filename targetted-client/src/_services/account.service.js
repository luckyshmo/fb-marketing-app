import store from '@src/store/store'

export const accountService = {
  login,
  logout,
  getInstagram,
}

function getInstagram (pageID) {
  FB.api(
    "/{page-id}/instagram_accounts",
    function (response) {

      if (response && !response.error) {

      }
    }
  );
}

async function login () {

  FB.login(function (authResponse) {
    if (!authResponse) return
    store.dispatch('setAccount', authResponse.authResponse.accessToken)

  }, {
    scope: 'ads_management, public_profile, business_management, pages_show_list, instagram_basic, instagram_content_publish, read_insights, instagram_manage_insights, pages_read_engagement, pages_manage_ads',
    return_scopes: true
  })
}

function logout () {
  // revoke app permissions to logout completely because FB.logout() doesn't remove FB cookie
  store.dispatch('setAccount', '')
  FB.api('/me/permissions', 'delete', null, () => FB.logout())
}

export default accountService
