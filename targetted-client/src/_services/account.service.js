// import { BehaviorSubject } from 'rxjs';

import store from '@src/store/store'

// const accountSubject = new BehaviorSubject(null);

export const accountService = {
  login,
  logout
}

async function login () {
  console.log('startLoginFB')
  FB.login(function (authResponse) {
    if (!authResponse) return
    console.log(authResponse.authResponse.accessToken)
    console.log(authResponse)
    store.dispatch('setAccount', authResponse.authResponse.accessToken)

    console.log('AccountSet')
  }, {
    scope: 'ads_management, public_profile, business_management, pages_show_list, instagram_basic, instagram_content_publish, read_insights, instagram_manage_insights, pages_read_engagement',
    // manage_pages, publish_pages
    return_scopes: true
  })
}

function logout () {
  // revoke app permissions to logout completely because FB.logout() doesn't remove FB cookie
  store.dispatch('setAccount', '')
  FB.api('/me/permissions', 'delete', null, () => FB.logout())
}

export default accountService
