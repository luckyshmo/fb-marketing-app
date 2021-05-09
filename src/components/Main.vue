<template>
    <div id="content">
        <h1 id="h1">Главная</h1>
        <h2 id="h2">Ваши рекламные кампании</h2>
        <p id="p1">Создайте рекламную кампанию, чтобы приводить новых клиентов в ваш бизнес</p>
        <router-link :to="{name: 'createCompany'}">
            <b-button variant="primary" id="main-button">Создать компанию</b-button>
        </router-link>
    </div>
</template>
<script>
export default {
    
}
</script>
<style>

</style>