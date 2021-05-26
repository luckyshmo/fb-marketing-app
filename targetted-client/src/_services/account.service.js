// import { BehaviorSubject } from 'rxjs';

import store from '../../store/store'

// const accountSubject = new BehaviorSubject(null);

export const accountService = {
    login,
    logout,
};

async function login() {
    console.log("startLoginFB")
    FB.login(function(authResponse) {
        if (!authResponse) return;
        console.log(authResponse.authResponse.accessToken)
        console.log(authResponse)
        store.dispatch('setAccount', authResponse.authResponse.accessToken);

        console.log("AccountSet")
    }, {
        scope: 'email, ads_management, instagram_basic, business_management, pages_show_list',
        return_scopes: true
    });
}

function logout() {
    // revoke app permissions to logout completely because FB.logout() doesn't remove FB cookie
    store.dispatch('setAccount', '');
    FB.api('/me/permissions', 'delete', null, () => FB.logout());
}

export default accountService;