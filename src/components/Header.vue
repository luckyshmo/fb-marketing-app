<template>
    <div id="container">
        <div class="nav" id="left">
            <p>Написать нам</p>
        </div>
        <div class="logo" id="middle">
            <p>Targetted</p>
        </div>
        <div class="nav" id="right">
            <p>Пользователь</p>
        </div>
    </div>
</template>
<script>
</script>
<style>
.logo {
    font-family: Monument Extended;
    font-style: normal;
    font-weight: normal;
    font-size: 28px;
    line-height: 100%;

    letter-spacing: 0.8em;
    text-transform: uppercase;

    background: -webkit-linear-gradient(left, #0704FD, #F498AD, #BB44CB);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;

    width: 50%;
}
.nav {
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    font-size: 16px;
    line-height: 24px;
    text-align: center;

    color: #FFFFFF;

    width: 25%;
}

#container {height: 100%; width:100%; text-align: center;}
#left, #middle, #right {display: inline-block; *display: inline; zoom: 1; vertical-align: center;}
#left {width: 25%;}
#middle {width: 50%;}
#right {width: 25%;}
    
</style>