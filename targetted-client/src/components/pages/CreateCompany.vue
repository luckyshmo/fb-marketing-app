<template>
    <div>
        <div id="content-wrapper">
            <div id="content">
            <router-link :to="{name: 'mainPage'}">
                <p id="navigation-text" style="margin:0;">← К списку кампаний</p>
            </router-link>
            <h1 id="h1">Создание кампании</h1>
            <div>
                <b-form @submit.prevent="createCompany()">
                    <h2 id="h2">Доступ к кабинету Facebook</h2>
                    <div v-if="!(store.getters.GET_FB_PAGES.length > 0) && !isRequestSent && !pageSubmitted">
                        <p id="p1">Привяжите свой аккаунт Facebook к targetted, чтобы натсроить и запустить рекламунюу компанию </p>
                        <b-button 
                            v-if="!(store.getters.GET_FB_PAGES.length > 0)"
                            variant="primary"
                            id="main-button"
                            @click="loginFB"
                        >
                            У меня есть бизнес-аккаунт
                        </b-button>
                        <popup
                        v-if="isInfoPopupVisible"
                        popupTitle="Инструкция по созданию бизнесс-аккаунта"
                        @closePopup="closeInfoPopup"
                        >
                            <div>
                            <p id="p1">Тип инструкция все дела</p>
                            </div>
                        </popup>
                        <b-button 
                            v-if="!(store.getters.GET_FB_PAGES.length > 0)"
                            variant="primary"
                            id="main-button"
                            style="margin-left: 15px;"
                            @click="showPopupInfo"
                        >
                            Нет бизнесс-аккаунта
                        </b-button>
                    </div>
                    <div v-if="store.getters.GET_FB_PAGES.length > 0 && !isRequestSent && !pageSubmitted">
                        <div>
                            <p id="p1">Выберите страницу которую хотите привязать</p>
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
                        <b-button 
                            variant="primary"
                            id="main-button"
                            @click="sendFbRequest()"
                        >
                            Привязать
                        </b-button>
                    </div>
                    <div v-if="isRequestSent && !pageSubmitted">
                        <p id="p1">Зайди в аккаунт на Facebook и подтверди привязку страницы в сообщениях</p>
                        <b-button 
                            variant="primary"
                            id="main-button"
                            target="_blank"
                            rel="noopener noreferrer"
                            href="https://facebook.com"
                        >
                        Перейти в facebook
                        </b-button>
                        <b-button 
                            variant="primary"
                            id="main-button"
                            style="margin-left: 15px"
                            @click="checkPageSubmitted()"
                        >
                            Я подтвердил в сообщениях
                        </b-button>
                    </div>
                    <div v-if="pageSubmitted">
                        <div class="c-status" style="margin-top: 30px;">
                            <div class="elipse" id="green"></div>
                            <p class="c-status-text" style="text-align: left;" 
                            >Страница {{form.fbPageId}} привязана к targetted</p>
                        </div>
                    </div>
                    <b-button 
                        v-if="store.getters.GET_FB_PAGES.length > 0"
                        variant="primary"
                        id="main-button"
                        style="margin-top: 30px; background: #F3F3F3; color: black"
                        @click="logout"
                    >
                        Привзять другой аккаунт
                    </b-button>
                    
                    <h2 id="h2">Основное</h2>

                    <b-form-group
                        label="Название компании"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                    >
                        <b-form-input
                        class="form-input"
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
                        class="form-input"
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
                        class="form-input"
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
                                    class="x-button"
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
                            class="form-input"
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
                                class="x-button"
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
                            class="form-input"
                            v-model="form.imagesSmallDescription[index]"
                            placeholder="Введите текст"
                            ></b-form-input>
                        </b-form-group>
                    </div>
                </div>
                <b-form-group
                v-if="form.imagesSmall.length > 0"
                        label="Описание под постом в ленте"
                        :label-cols="label_cols"
                        :content-cols="content_cols"
                        id="input-group1"
                        label-for="input-horizontal"
                    >
                        <b-form-textarea
                        class="form-input"
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
// import axios from 'axios'
import popup from '../popup.vue'
export default {
    name: 'CreateCompany',
    components: {
      popup
    },
    data() {
        return{
            store,
            isInfoPopupVisible: false,
            label_cols: this.getWidth().label,
            content_cols: this.getWidth().content,
            isRequestSent: false,
            pageSubmitted: false,
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
        getWidth() {
            let width = Math.max(
                document.body.scrollWidth,
                document.documentElement.scrollWidth,
                document.body.offsetWidth,
                document.documentElement.offsetWidth,
                document.documentElement.clientWidth
            );
            if (width < 570){
                return {
                    label:12,
                    content:12
                };
            }
            return {
                    label:3,
                    content:9
            };
        },
        closeInfoPopup() {
            this.isInfoPopupVisible = false;
        },
        showPopupInfo() {
            this.isInfoPopupVisible = true;
        },
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
            // sendFbRequest() //TODO
            console.log("Page ID", this.form.fbPageId)
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
                companyData.append("imageSmall", image); //TODO не прилетают.
            });
            Array.from(this.form.images).forEach(image => {
                companyData.append("image", image);
            });
            store.dispatch("saveCompany", companyData)
            .then((resp)=>{
                console.log(resp.data) //TODO log company id
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
                this.isInfoPopupVisible = false
                this.isRequestSent = false
                router.push('main')
            })
            .catch(err => {
                console.log(err) //TODO popup
            })
        },
        sendFbRequest(){
            // let appToken = 'EAAEDuTXOcAgBAEbAJLLg00LDOJH4LyOekYZCWtJhjul3xbrUpQZCWt0LEDTlpQrsxhwWUZBSjZAA5OyRMgZB0g83zIIKXNQRys82ZAajuUGAmZAmQGy5kH242uZAZABoMjgebiuGQkcjKJ5Kd8xyWXThFQytJP1ATmHNNQvPZA0I1RROQAbmWUJS8HgyFMtWkETMecbEPUNLC4zgZDZD'
            // let companyId = 856950044859235
            // let files = {
            //     'page_id': this.form.fbPageId,
            //     'permitted_tasks': '[\'MANAGE\', \'CREATE_CONTENT\', \'MODERATE\', \'ADVERTISE\', \'ANALYZE\']',
            //     'access_token': appToken,
            // }
            // axios.post(`https://graph.facebook.com/v10.0/${companyId}/client_pages`, files)
            // .then(
            //     this.isRequestSent = true
            // )
            // .catch(err => {
            //     console.log(err) //TODO popup
            // })

            this.isRequestSent = true
        },
        checkPageSubmitted(){
             //         	curl -G \
	// -d "access_token=EAAEDuTXOcAgBAEbAJLLg00LDOJH4LyOekYZCWtJhjul3xbrUpQZCWt0LEDTlpQrsxhwWUZBSjZAA5OyRMgZB0g83zIIKXNQRys82ZAajuUGAmZAmQGy5kH242uZAZABoMjgebiuGQkcjKJ5Kd8xyWXThFQytJP1ATmHNNQvPZA0I1RROQAbmWUJS8HgyFMtWkETMecbEPUNLC4zgZDZD" \
	// "https://graph.facebook.com/v10.0/856950044859235/client_pages"
            this.pageSubmitted = true
        },
        logout(){
            this.pageSubmitted = false
            this.isRequestSent = false
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