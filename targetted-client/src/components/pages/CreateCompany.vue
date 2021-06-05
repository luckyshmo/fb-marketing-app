<template>
    <div>
        <div id="content-wrapper">
            <div id="content">
            <router-link :to="{name: 'mainPage'}">
                <p id="navigation-text" style="margin:0;">← К списку кампаний</p>
            </router-link>
            <h1 id="h1">Создание компании</h1>
            <h2 id="h2">Доступ к кабинету Facebook</h2>
            <p id="p1">Привяжите свой аккаунт Facebook к targetted, чтобы натсроить и запустить рекламунюу компанию </p>
            <b-button 
                v-if="!(store.getters.GET_FB_PAGES.length > 0)"
                variant="primary"
                id="main-button"
                @click="loginFB"
            >
                Настроить доступ
            </b-button>
            <b-button 
                v-if="store.getters.GET_FB_PAGES.length > 0"
                variant="primary"
                id="main-button"
                @click="logout"
            >
                Разлогин
            </b-button>
            <div>
                <b-form @submit.prevent="createCompany()">
                    <div v-if="store.getters.GET_FB_PAGES.length > 0">
                        <h2 id="h2">Выбор страницы</h2>
                        <b-form-group
                        label="Выберите страницу"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                        >
                            <b-form-radio-group
                                v-model="form.fbPageId"
                                :options="store.getters.GET_FB_PAGES"
                            ></b-form-radio-group>
                        </b-form-group>
                    </div>

                    <h2 id="h2">Основное</h2>

                    <b-form-group
                        label="Название компании"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                    >
                        <b-form-input
                        id="form-input"
                        required
                        v-model="form.companyName"
                        placeholder="Введите название"
                        ></b-form-input>
                    </b-form-group>

                    <b-form-group
                        label="Цель кампании"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                    >
                        <b-form-radio-group
                            v-model="form.compnayPurpose"
                            :options="[
                                'Сообщения в директ',
                                'Лиды через лидформу',
                                'Конверсия на сайте'
                            ]"
                        ></b-form-radio-group>
                    </b-form-group>

                    <b-form-group
                        label="Сфера деятельности"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                    >
                        <b-form-input
                        id="form-input"
                        v-model="form.companyField"
                        placeholder="Введите сферу"
                        ></b-form-input>
                    </b-form-group>

                    <b-form-group
                        label="Адрес бизнеса"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                        description="Нужен только для офлайн бизнеса"
                    >
                        <b-form-input
                        id="form-input"
                        v-model="form.businessAdress"
                        placeholder="Точный адрес"
                        ></b-form-input>
                    </b-form-group>
                <h2 id="h2">Креативы</h2>
                <b-form-group
                        label="Наличие креативов"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                >
                    <b-form-radio-group
                        v-model="form.creativeStatus"
                        :options="[
                            'Есть рекламные креативы',
                            'Создать рекламные креативы'
                        ]"
                    ></b-form-radio-group>
                </b-form-group>
                <b-form-group
                        :label="getStoriesLabel()"
                        :label-cols="label_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                        description="До 5 слайдов в сториз"
                >
                    <div id="block">
                        <div 
                        v-for="(image, key) in form.images" 
                        :key="key">
                            <div id="image-preview">
                                <div 
                                style="position: absolute;
                                margin-left: 140px;
                                margin-top: -15px;">
                                    <b-icon
                                    @click="removeImage(image)"
                                    style="height: 35px;
                                            width: 35px;
                                            background: #e4e4e4;
                                            border-radius: 17.5px;"
                                    icon="x"></b-icon>
                                </div>
                                <img id="preview" :ref="'image'" />
                            </div>
                        </div>
                        
                        <div v-if="this.form.images.length < 5">
                            <input 
                            style="display: none"
                            type="file" 
                            multiple
                            accept="image/gif, image/jpeg, image/png, image/jpg" 
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
                </b-form-group>
                <div v-if="isCreative()">
                    <div 
                    v-for="(image, index) in form.images" 
                    :key="image.name">
                        <b-form-group
                            :label="textOnSlide(index)"
                            :label-cols="label_cols"
                            :content-cols="content_cols"
                            id="input-group1"
                            label-for="input-horizontal"
                        >
                            <b-form-input
                            id="form-input"
                            v-model="form.imagesDescription[index]"
                            placeholder="Введите текст"
                            ></b-form-input>
                        </b-form-group>
                    </div>
                </div>
                <b-form-group
                        :label="getPostLabel()"
                        :label-cols="label_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                        description="До 5 слайдов в посте"
                >

                    <div id="block">
                    <div 
                    v-for="(image, key) in form.imagesSmall" 
                    :key="key">
                        <div id="image-preview">
                            <div 
                            style="position: absolute;
                            margin-left: 140px;
                            margin-top: -15px;">
                                <b-icon
                                @click="removeImageSmall(image)"
                                style="height: 35px;
                                        width: 35px;
                                        background: #e4e4e4;
                                        border-radius: 17.5px;"
                                icon="x"></b-icon>
                            </div>
                            <img id="preview-small" :ref="'imageSmall'" />
                        </div>
                    </div>
                    <div v-if="this.form.imagesSmall.length < 5">
                        <input 
                        style="display: none"
                        type="file" 
                        multiple
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
                    
                </b-form-group>
                <div v-if="isCreative()">
                    <div 
                    v-for="(image, index) in form.imagesSmall" 
                    :key="image.name">
                        <b-form-group
                            :label="textOnImage(index)"
                            :label-cols="label_cols"
                            :content-cols="content_cols"
                            id="input-group1"
                            label-for="input-horizontal"
                        >
                            <b-form-input
                            id="form-input"
                            v-model="form.imagesSmallDescription[index]"
                            placeholder="Введите текст"
                            ></b-form-input>
                        </b-form-group>
                    </div>
                </div>
                <b-form-group
                v-if="isCreative()"
                        label="Описание под постом в ленте"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                    >
                        <b-form-textarea
                        id="form-input"
                        style="height: 100px"
                        v-model="form.postDescription"
                        placeholder="Введите текст"
                        ></b-form-textarea>
                    </b-form-group>
                    <b-button
                        id="submit-button"
                        type="submit"
                    >
                        Создать компанию
                    </b-button>
                </b-form>
            </div>
        </div>
        </div>
    </div>
    
</template>
<script>
import accountService from '../../_services/account.service';
import store from '../../../store/store'
import router from '../../../router/router'
import axios from 'axios'
export default {
    name: 'CreateCompany',
    data() {
        return{
            store,
            label_cols: 3,
            content_cols: 9,
            form: {
                fbPageId: '',
                companyName: '',
                compnayPurpose: '',
                companyField: '',
                businessAdress: '',
                images: [],
                imagesDescription: [],
                imagesSmall: [],
                imagesSmallDescription: [],
                creativeStatus: '',
                postDescription: '',
            },
        }
    },
    methods: {
        removeImageSmall(image){
            this.remove(this.form.imagesSmall, image);
            for (let i = 0; i < this.form.imagesSmall.length; i++) {
                let reader = new FileReader();
                reader.onload = () => {
                    this.$refs.imageSmall[i].src = reader.result;
                };

                reader.readAsDataURL(this.form.imagesSmall[i]);
            }
        },
        removeImage(image){
            this.remove(this.form.images, image)
            for (let i = 0; i < this.form.images.length; i++) {
                let reader = new FileReader();
                reader.onload = () => {
                    this.$refs.image[i].src = reader.result;
                };

                reader.readAsDataURL(this.form.images[i]);
            }
        },
        remove(arr, img) {
            for (let i = 0; i < arr.length; i++) {
                if (arr[i].name === img.name){
                    arr.splice(i, 1)
                }
            }
        },
        getStoriesLabel(){
            if (!this.isCreative()) {
                return "Креативы для Сториз"
            }
            return "Картинки для Сториз"
        },
        getPostLabel(){
            if (!this.isCreative()) {
                return "Креативы для поста в ленте"
            }
            return "Картинки для поста в ленте"
        },
        isCreative(){
            return this.form.creativeStatus === 'Создать рекламные креативы'
        },
        textOnSlide(index){
            return `Текст на слайде ${index+1}`
        },
        textOnImage(index){
            return `Текст на картинке ${index+1}`
        },
        loginFB() {
            accountService.login()
        },
        createCompany(){
            // sendFbRequest()
            let api = process.env.VUE_APP_API_URL
            console.log("Api adrr:", api)
            const companyData = new FormData();
            companyData.append("fbPageId", this.form.fbPageId)
            companyData.append("companyName", this.form.companyName)
            companyData.append("compnayPurpose", this.form.compnayPurpose)
            companyData.append("companyField", this.form.companyField)
            companyData.append("businessAdress", this.form.businessAdress)
            companyData.append("imagesDescription", this.form.imagesDescription)
            companyData.append("imagesSmallDescription", this.form.imagesSmallDescription)
            companyData.append("creativeStatus", this.form.creativeStatus)
            companyData.append("postDescription", this.form.postDescription)
            Array.from(this.form.imagesSmall).forEach(image => {
                companyData.append("imageSmall", image);
            });
            Array.from(this.form.images).forEach(image => {
                companyData.append("image", image);
            });
            store.dispatch("saveCompany", companyData)
            this.form = {
                fbPageId: '',
                companyName: '',
                compnayPurpose: '',
                companyField: '',
                businessAdress: '',
                images: [],
                imagesDescription: [],
                imagesSmall: [],
                imagesSmallDescription: [],
                creativeStatus: '',
                postDescription: '',
            }
            router.push('main')
        },
        sendFbRequest(){
            let appToken = 'EAAEDuTXOcAgBAEbAJLLg00LDOJH4LyOekYZCWtJhjul3xbrUpQZCWt0LEDTlpQrsxhwWUZBSjZAA5OyRMgZB0g83zIIKXNQRys82ZAajuUGAmZAmQGy5kH242uZAZABoMjgebiuGQkcjKJ5Kd8xyWXThFQytJP1ATmHNNQvPZA0I1RROQAbmWUJS8HgyFMtWkETMecbEPUNLC4zgZDZD'
            let companyId = 856950044859235
            let files = {
                'page_id': this.form.fbPageId,
                'permitted_tasks': '[\'MANAGE\', \'CREATE_CONTENT\', \'MODERATE\', \'ADVERTISE\', \'ANALYZE\']',
                'access_token': appToken,
            }
            axios.post(`https://graph.facebook.com/v10.0/${companyId}/client_pages`, files)
        },
        logout(){
            accountService.logout()
        },
        onFileSelected(e) {
            let selectedFiles = e.target.files;
            let len = Math.min(selectedFiles.length, 5)
            for (let i = 0; i < len; i++) {
                this.form.images.push(selectedFiles[i]);
            }

            for (let i = 0; i < this.form.images.length; i++) {
                let reader = new FileReader();
                reader.onload = () => {
                    this.$refs.image[i].src = reader.result;
                };

                reader.readAsDataURL(this.form.images[i]);
            }
        },
        onSmallFileSelected(e) {
            let selectedFiles = e.target.files;
            let len = Math.min(selectedFiles.length, 5)
            for (let i = 0; i < len; i++) {
                this.form.imagesSmall.push(selectedFiles[i]);
            }

            for (let i = 0; i < this.form.imagesSmall.length; i++) {
                let reader = new FileReader();
                reader.onload = () => {
                    this.$refs.imageSmall[i].src = reader.result;
                };

                reader.readAsDataURL(this.form.imagesSmall[i]);
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
.custom-radio{
    margin: 10px;
}
.custom-control-input{
    margin-right: 3px;
}
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
    grid-template-columns: repeat(auto-fill,minmax(200px, 1fr));
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
#list {
  font-family: Montserrat;
  font-style: normal;
  font-weight: normal;
  font-size: 20px;
  line-height: 32px;

  color: #000000;

}
#list-item {
    border-radius: 8px;
    border: none;
    background-color: #e4e4e4;
    margin: 5px;

}
</style>