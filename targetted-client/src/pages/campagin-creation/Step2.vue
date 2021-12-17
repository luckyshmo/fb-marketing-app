<template>
  <div class="step step1">
    <slot name="header"></slot>
    <p>
      Загрузите картинки, которые будут отображаться в рекламе.
    </p>
    <b-form>
      <div v-if="isEdit">
        <b-form-group v-if="imageNames.length > 0" label="Уже имеющиеся креативы" :label-cols="label_cols"
          :content-cols="content_cols" id="input-group-main" label-for="input-horizontal">
          <div id="image-block">
            <div v-for="(name) in imageNames" :key="name">
              <div>
                <img id="preview" :src="getImageByName(name)" />
              </div>
            </div>
          </div>
        </b-form-group>
      </div>

      <div v-if="!isEdit">
        <!-- <b-form-group label="Наличие креативов" :label-cols="label_cols" :content-cols="content_cols"
          id="input-group-main" label-for="input-horizontal">
          <b-form-radio-group
          v-model="campaign.CreativeStatus"
          class="app-new-radio"
          :options="[
                                  'Есть рекламные креативы',
                                  'Создать рекламные креативы'
                              ]"></b-form-radio-group>
        </b-form-group> -->

        <div class="creative-message">
           <b-row>
              <b-col cols="10">
                Советы по самостоятельному созданию креативов
              </b-col>
              <b-col cols="2" class="right-pointer">
                <b-icon icon="chevron-right"/>
              </b-col>
           </b-row>

          <!-- <div v-if="isCreative()">
            Для создания рекламных креативов загрузите картинки и напишите текст, который будет на них отображаться. <a
              href="https://docs.google.com/document/d/1gqOkpxDJ1wNt-AYlt5Q1Et1kF8NRLRYdG-dXK7WdT1k/edit?usp=sharing"
              style="color: #6C1BD2" target="_blank">Посмотрите пример,</a> как это будет выглядеть и подробную
            инструкцию.
          </div>
          <div v-if="!isCreative()">
            <a href="https://docs.google.com/document/d/1DGJQw2lw_Y6KzyPeqA6To8u9UuiYi2Kokx_yjvTxgKE/edit?usp=sharing"
              style="color: #6C1BD2" target="_blank">Воспользуйтесь советами</a> при самостоятельном создании креативов,
            чтобы увеличить эффективность рекламной кампании.
          </div> -->
        </div>
      <h2 class="app-header-bold">{{getStoriesLabel()}}</h2>

      <p class="app-new-info">В одном рекламном объявлении может быть до 5 слайдов</p>

        <b-form-group :label-cols="label_cols" :state="validateImages()" id="input-group-main"
         class="app-new-creative-uploads"
          label-for="input-horizontal">
          <div id="image-block">
            <div v-if="Images.length < 5" key="s0">
              <input style="display: none" type="file" multiple accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                @change="onFileSelected" ref="fileInput">
              <div @click="$refs.fileInput.click()" id="load-frame">
                <!-- Обводка у кнопок -->
                <!-- ПОехал размер -->
                <p id="load-file">Загрузить<br>фото или видео</p>
                <p id="file-size-big">Размер<br>1920х1080рх</p>
                <!-- TODO FILL -->
              </div>
            </div>

            <div v-for="(Image, key) in Images" :key="key">
              <div class="stories-image-wrapper">
                <div class="icon-div-image">
                  <b-icon @click.stop="removeImage(Image)" class="x-button" icon="x"></b-icon>
                </div>
                <div class="icon-div-number">{{key+1}}</div>
                <img id="preview" :ref="'Image'"  alt=""/>
              </div>
            </div>

          </div>
          <b-form-invalid-feedback class="error-message">
            Необходим минимум один файл
          </b-form-invalid-feedback>
        </b-form-group>
        <div v-if="isCreative()">
          <div v-for="(Image, index) in Images" :key="Image.name">
            <b-form-group :label="textOnSlide(index)" :label-cols="label_cols" :content-cols="content_cols"
              id="input-group-main" label-for="input-horizontal">
              <b-form-input class="form-input" v-model="campaign.ImagesDescription[index]" placeholder="Введите текст">
              </b-form-input>
            </b-form-group>
          </div>
        </div>
        <h2 class="app-header-bold">{{getPostLabel()}}</h2>
        <p class="app-new-info">В одном рекламном объявлении может быть до 5 слайдов</p>

        <b-form-group class="app-new-creative-uploads" :label-cols="label_cols" :state="validateImagesSmall()"
          id="input-group-main" label-for="input-horizontal">

          <div id="image-small-block">
            <div v-if="ImagesSmall.length < 5" key="s0">
              <input style="display: none" type="file" multiple accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                @change="onSmallFileSelected" ref="smallFileInput">
              <div @click="$refs.smallFileInput.click()" id="load-frame-small">
                <p id="load-file">Загрузить<br>фото или видео</p>
                <p id="file-size">Размер<br>1080х1080рх</p>
              </div>
            </div>

            <div v-for="(Image, key) in ImagesSmall" :key="key">
              <div class="post-image-wrapper">
              <div class="icon-div-image">
                <b-icon @click.stop="removeImageSmall(Image)" class="x-button" icon="x"></b-icon>
              </div>
              <div class="icon-div-number">{{key+1}}</div>
              <img id="preview-small" :ref="'ImageSmall'"  alt=""/>
              </div>
            </div>
          </div>
          <b-form-invalid-feedback class="error-message">
            Необходим минимум один файл
          </b-form-invalid-feedback>
        </b-form-group>
        <div v-if="isCreative()" class="app-new-creative-uploads">
          <div v-for="(Image, index) in ImagesSmall" :key="Image.name">
            <b-form-group :label="textOnImage(index)" :label-cols="label_cols" :content-cols="content_cols"
              id="input-group-main" label-for="input-horizontal">
              <b-form-input class="form-input" v-model="campaign.ImagesSmallDescription[index]"
                placeholder="Введите текст"></b-form-input>
            </b-form-group>
          </div>
        </div>
      </div>
      <b-form-group v-if="ImagesSmall.length > 0" label="Описание под постом в ленте" :label-cols="label_cols"
        :content-cols="content_cols" id="input-group-main" label-for="input-horizontal">
        <b-form-textarea class="form-input" style="height: 100px" v-model="campaign.PostDescription"
          placeholder="Введите текст"></b-form-textarea>
      </b-form-group>
        <h2>Предпросмотр</h2>
        <div class="creative-message" @click="$router.push('/preview')">
        <b-row>
              <b-col cols="10">
                      Как будет выглядеть реклама
              </b-col>
              <b-col cols="2" class="right-pointer">
                <b-icon icon="chevron-right"/>
              </b-col>
           </b-row>

        </div>

       <b-row class="mt-5 bottom__block">
        <b-col cols="6" sm="4" md="4" lg="4" xl="4">
          <b-button  type="button"
                    class="app-new-submit-button"
                    :class="{'disabled-look': validateImagesSmall() && validateImages()}"
                    @click="sendData">
            {{isEdit ? "Назад":"Продолжить"}}
          </b-button>
        </b-col>
        <b-col cols="6" sm="8" md="8" lg="8" xl="8" class="app-new-small-text-fit">
            <p class="text-muted text-pt-desktop">Загрузите креативы,<br/> чтобы продолжить</p>
        </b-col>
      </b-row>

    </b-form>
  </div>
</template>

<script>
import store from '@src/store/store'
import { APP_UI_URL } from '@src/constants'

export default {
  name: 'Step2',
  props: ['label_cols', 'content_cols', 'campaign', 'isEdit'],
  data: function () {
    return {
      store, // fixme
      imageNames: [],
      ImagesSmall: [],
      Images: []
    }
  },
  methods: {
    sendData () {
      if (this.validateImagesSmall() && this.validateImages()) {
        return
      }
      this.$router.push('/campaign-step-3')
      // this.$emit('next', {
      //   imageNames: this.imageNames,
      //   ImagesSmall: this.ImagesSmall,
      //   Images: this.Images
      // })
    },
    getImageByName (name) {
      const uID = this.campaign.UserId
      const cID = this.campaign.Id
      return `${APP_UI_URL}images/${uID}/${cID}${name}`
    },
    removeImageSmall (Image) {
      this.remove(this.ImagesSmall, Image)
      for (let i = 0; i < this.ImagesSmall.length; i++) {
        const reader = new FileReader()
        reader.onload = () => {
          this.$refs.ImageSmall[i].src = reader.result
        }

        reader.readAsDataURL(this.ImagesSmall[i])
      }
    },
    getPostLabel () {
      if (!this.isCreative()) {
        return 'Для поста в ленте'
      }
      return 'Для поста в ленте'
    },
    getStoriesLabel () {
      if (!this.isCreative()) {
        return 'Для сториз'
      }
      return 'Для сториз'
    },

    isCreative () {
      return this.campaign.CreativeStatus === 'Создать рекламные креативы'
    },
    validateImagesSmall () {
      return this.ImagesSmall.length === 0
    },
    validateImages () {
      return this.Images.length === 0
    },
    onFileSelected (e) {
      if (e != null) {
        const selectedFiles = e.target.files
        const len = Math.min(selectedFiles.length, 5 - this.Images.length)
        for (let i = 0; i < len; i++) {
          this.Images.push(selectedFiles[i])
        }

        for (let i = 0; i < this.Images.length; i++) {
          const reader = new FileReader()
          reader.onload = () => {
            this.$refs.Image[i].src = reader.result
          }

          reader.readAsDataURL(this.Images[i])
        }
      }
    },
    onSmallFileSelected (e) {
      if (e != null) {
        const selectedFiles = e.target.files
        const len = Math.min(selectedFiles.length, 5 - this.ImagesSmall.length)
        for (let i = 0; i < len; i++) {
          this.ImagesSmall.push(selectedFiles[i])
        }

        for (let i = 0; i < this.ImagesSmall.length; i++) {
          const reader = new FileReader()
          reader.onload = () => {
            this.$refs.ImageSmall[i].src = reader.result
          }

          reader.readAsDataURL(this.ImagesSmall[i])
        }
      }
    },
    remove (arr, img) {
      for (let i = 0; i < arr.length; i++) {
        if (arr[i].name === img.name) {
          arr.splice(i, 1)
        }
      }
    },
    removeImage (Image) {
      this.remove(this.Images, Image)
      for (let i = 0; i < this.Images.length; i++) {
        const reader = new FileReader()
        reader.onload = () => {
          this.$refs.Image[i].src = reader.result
        }

        reader.readAsDataURL(this.Images[i])
      }
    },
    textOnSlide (index) {
      return `Текст на слайде ${index + 1}`
    },
    textOnImage (index) {
      return `Текст на картинке ${index + 1}`
    }
  }
}
</script>

<style lang="scss">
  @import '@src/assets/styles/vars.scss';
  .bottom__block{
    margin-bottom: 200px;
  }
  .app-new-creative-uploads {
    label[for=input-horizontal] {
      display: none;
    }
  }
  #input-group-main{
    margin-bottom: 40px !important;
  }
  #image-small-block{
    width: 100%;
    display: grid;
    grid-template-columns: repeat(3, minmax(150px, 180px));
    justify-content: space-between;
    align-items: center;
    grid-gap: 20px;
    &  > div{
      margin-top: 20px;
      width: 180px;
      height: 180px;
    }
  }
  #preview{
      width: 100%;
      height: 100%;
      object-fit: cover;
      border-radius: 12px;
  }

  #preview-small{
      width:100%;
      height:100%;
      object-fit: cover;
      border-radius: 20px;
  }

  #image-block {
      width: 100%;
      display: grid;
      grid-template-columns: repeat(3,minmax(150px, 180px));
      justify-content: space-between;
      align-items: center;
      grid-gap: 20px;
      //overflow-x: scroll;
      //margin-top: -30px;
    & > div{
      max-width: 180px;
      height: 316px;
      margin-top: 20px;
    }
  }

  #load-file {
      font-style: normal;
      font-weight: 600;
      font-size: $baseFont;
      line-height: $baseLH;
      margin: 28px auto 18px;
      text-align: center;
      color: $black;
  }

  #file-size-big {
      font-style: normal;
      font-weight: normal;
      font-size: 12px;
      line-height: 20px;
      /* bottom: 0;//TODO почему то не работает */
      margin-top: 134px;
      text-align: center;
      color: $gray;
  }

  #file-size {
      font-style: normal;
      font-weight: normal;
      font-size: 12px;
      line-height: 20px;
      bottom: 0;
      text-align: center;
      color: $gray;
  }
  .stories-image-wrapper {
    width: 100%;
    height: 100%;
    position: relative;
  }
  .icon-div-image{
    position: absolute;
    top: -10px;
    right: -10px;
  }

  //@media (max-width: 600px) {
  //  .step{
  //    width: 100%;
  //  }
  //}
  //@media (max-width: 450px) {
  //    #image-block {
  //        max-width: 200px;
  //    }
  //}

  #load-frame {
      border: 2px dashed #CCCCCC;
      border-radius: 16px;
      //width: 174px;
      //height: 280px;
    width: 100%;
    height: 100%;
  }

  #load-frame-small {
      border: 2px dashed #CCCCCC;
      border-radius: 16px;
      width: 180px;
      height: 180px;
  }
  .post-image-wrapper {
    position: relative;
    width: 100%;
    height: 100%;
  }
  @media (min-width: 650px) {
    //#load-frame {
    //  width: 180px;
    //  height: 316px;
    //}
    //#load-frame-small {
    //  width: 201px;
    //  height: 201px;
    //}
  }

  #load-frame:hover, #load-frame-small:hover {
      background: $light;
  }

  .post-image-wrapper {
    //margin-top: -30px;
  }
  .icon-div-number {
    bottom: 10px;
    position: absolute;
    background-color: $black;
    border-radius: 12px;
    color: $light;
    width: 32px;
    height: 32px;
    font-size: 14px;
    margin-bottom: 0;
    padding: 6px 13px;
    left: 10px;
  }

  .post-image-wrapper {
    //margin-top: -30px;

    & .icon-div-number {
      //bottom: -150px;
    }
  }

</style>
