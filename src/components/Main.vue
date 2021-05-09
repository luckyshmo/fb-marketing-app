<template>
    <div id="content">
        <h1 id="h1">Главная</h1>
        <h2 id="h2">Ваши рекламные кампании</h2>
        <p id="p1">Создайте рекламную кампанию, чтобы приводить новых клиентов в ваш бизнес</p>
        <b-button variant="primary" id="main-button">Создать компанию</b-button>
    </div>
</template>
<script>
export default {
    
}
</script>
<style>
#main-button {
    color:whitesmoke;
    background-color: #6C1BD2;
    border-radius: 8px;
    padding: 12px 28px;
    border:none;
    font-family: Montserrat;
    font-style: normal;
    font-weight: 600;
}

#content {
    padding: 40px;
}

#h1 {
  margin-top: 0px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 66px;
  line-height: 100%;

  color: #000000;
}
#p1 {
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-size: 20px;
  line-height: 32px;

  color: #000000;

}

#h2 {
    margin-top: 50px;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 900;
  font-size: 36px;
  line-height: 44px;

  color: #000000;
}

</style>