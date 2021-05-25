<template>
    <div id="content">
        <router-link :to="{name: 'mainPage'}">
            <p id="navigation-text" style="margin:0;">← К списку кампаний</p>
        </router-link>
        <h1 id="h1">Создание компании</h1>
        <h2 id="h2">Доступ к кабинету Facebook</h2>
        <p id="p1">Привяжите свой аккаунт Facebook к targetted, чтобы натсроить и запустить рекламунюу компанию </p>
        <b-button 
            variant="primary"
            id="main-button"
            @click="loginFB"
        >
            Настроить доступ
        </b-button>
        <b-button 
            variant="primary"
            id="main-button"
            @click="logout"
        >
            Разлогин
        </b-button>
        <!-- <div v-if="store.getters.GET_FB_TOKEN.len"> -->
            <p>{{store.getters.GET_FB_ACCOUNT}}</p>
        <!-- </div> -->
        <h2 id="h2">Основное</h2>
        
        <h2 id="h2">Креативы</h2>

        <div id="block">
            <!-- v-if="imagesData.length > 0" -->
            <div 
            v-for="(image, key) in images" 
            :key="key">
                <div id="image-preview">
                    <img id="preview" :ref="'image'" />
                </div>
            </div>
            
            <div>
                <input 
                style="display: none"
                type="file" 
                multiple
                @change="onFileSelected"
                ref="fileInput">
                <div 
                @click="$refs.fileInput.click()"
                id="load-frame">
                    <p id="load-file">Загрузить<br>файл</p>
                    <p id="file-size" style="margin: 140px auto;">Размер<br>1920х1080рх</p>
                </div>
            </div>
        </div>
        
        <div id="block">
            <div 
            v-for="(image, key) in imagesSmall" 
            :key="key">
                <div id="image-preview">
                    <img id="preview-small" :ref="'imageSmall'" />
                </div>
            </div>

            <input 
            style="display: none"
            type="file" 
            accept="image/gif, image/jpeg, image/png, image/jpg" 
            @change="onSmallFileSelected"
            ref="smallFileInput">
            <div 
            @click="$refs.smallFileInput.click()"
            id="load-frame-small">
                <p id="load-file">Загрузить<br>файл</p>
                <p id="file-size">Размер<br>1080х1080рх</p>
            </div>
        </div>
    </div>
</template>
<script>
import accountService from '../../_services/account.service';
import store from '../../../store/store'
import axios from 'axios'
export default {
    name: 'CreateCompany',
    data() {
        return{
            store,
            images: [],
            imagesSmall: []
        }
    },
    methods: {
        loginFB() {
            accountService.login()
        },
        logout(){
            accountService.logout()
        },
        onFileSelected(e) {
            let selectedFiles = e.target.files;
            for (let i = 0; i < selectedFiles.length; i++) {
                this.images.push(selectedFiles[i]);
            }

            for (let i = 0; i < this.images.length; i++) {
                let reader = new FileReader();
                reader.onload = () => {
                    this.$refs.image[i].src = reader.result;
                };

                reader.readAsDataURL(this.images[i]);
            }
        },
        onSmallFileSelected(e) {
            let selectedFiles = e.target.files;
            for (let i = 0; i < selectedFiles.length; i++) {
                this.imagesSmall.push(selectedFiles[i]);
            }

            for (let i = 0; i < this.imagesSmall.length; i++) {
                let reader = new FileReader();
                reader.onload = () => {
                    this.$refs.imageSmall[i].src = reader.result;
                };

                reader.readAsDataURL(this.imagesSmall[i]);
            }
        },
        onUpload() {
            const fd = new FormData();
            fd.append('image', this.selectedFiles[0], this.selectedFiles[0].name)
            axios.post('http://localhost:3000/api/file', fd, {
                onUploadProgress: uploadEvent => {
                    console.log("Upload Progress: " + Math.round(uploadEvent.loaded / uploadEvent.total * 100) + "%")
                }
            })
            .then(res => {
                console.log(res)
            })
        }
    }
}
</script>
<style>
#preview{
    width: 160px;
    height: 280px;
    border-radius: 20px;
}
#preview-small{
    width: 160px;
    height: 160px;
    border-radius: 20px;
}
#block {
    margin-top: 20px;
    display: grid;
    grid-template-columns: repeat(auto-fill,minmax(160px, 1fr));
    justify-content: space-between;
    align-items: center;
    grid-gap: 30px;
}
#load-frame {
    border: 2px dashed #CCCCCC;
    border-radius: 20px;
    width: 160px;
    height: 280px;
}
#load-frame-small {
    border: 2px dashed #CCCCCC;
    border-radius: 20px;
    width: 160px;
    height: 160px;
}
#load-file {
    font-family: Montserrat;
    font-style: normal;
    font-weight: 600;
    font-size: 16px;
    line-height: 24px;
    margin: 20px auto;
    text-align: center;
    color: #000000;
}
#file-size {
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    font-size: 12px;
    line-height: 20px;
    bottom: 0;
    text-align: center;
    color: #767676;
}
</style>