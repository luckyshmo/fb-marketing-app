<template>
  <div class="step">
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
                <img id="preview" :src="getImageByName(name)"/>
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

        <div class="creative-message" @click="showModalPopup()">
          <div>
            <b-col cols="10">
              Советы по самостоятельному созданию креативов
            </b-col>
            <b-col cols="2" class="right-pointer">
              <!--                <b-icon icon="chevron-right"/>-->
              <svg width="7" height="12" viewBox="0 0 7 12" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M1 1L5.5 6L1 11" stroke="black" stroke-width="2" stroke-linecap="round"/>
              </svg>
            </b-col>
          </div>

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
        <h2 class="app-header-bold">{{ getStoriesLabel() }}</h2>

        <p class="app-new-info">
          Лучше загрузить разные варианты
          <br class="display__none"/>рекламы, это повлияет <br class="display__block"/>на эффективность <br
          class="display__none"/>рекламы. <a href="#" class="app-new-action-text">Подробнее</a></p>

        <b-form-group :label-cols="label_cols" :state="validateImages()" id="input-group-main"
                      class="app-new-creative-uploads"
                      label-for="input-horizontal">
          <div id="image-block">
            <div v-for="(Image, index) in campaign.Images" :key="'key' + index" :style="imageBlockStyle()">
              <div class="stories-image-wrapper">
                <div class="icon-div-image">
                  <b-icon @click.stop="removeImage(Image , index)" class="x-button" icon="x"></b-icon>
                </div>
                <div class="icon-div-number">{{ index + 1 }}</div>
                <img id="preview" :src="campaign.ImagesSrc[index]" alt="img"/>
              </div>
            </div>
            <div v-if="campaign.Images.length < 5" key="s0" :style="imageBlockStyle()">
              <input style="display: none" type="file" multiple accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                     @change="onFileSelected" ref="fileInput">
              <div @click="$refs.fileInput.click()" id="load-frame">
                <!-- Обводка у кнопок -->
                <!-- ПОехал размер -->
                <h5 id="load-file">Загрузить<br>фото или видео</h5>
                <p id="file-size-big">Размер<br>1920х1080рх</p>
                <!-- TODO FILL -->
              </div>
            </div>

          </div>
          <b-form-invalid-feedback class="error-message">
            Необходим минимум один файл
          </b-form-invalid-feedback>
        </b-form-group>
        <!--        <div v-if="isCreative()">-->
        <!--          <div v-for="(Image, index) in Images" :key="Image.name">-->
        <!--            <b-form-group :label="textOnSlide(index)" :label-cols="label_cols" :content-cols="content_cols"-->
        <!--              id="input-group-main" label-for="input-horizontal">-->
        <!--              <b-form-input class="form-input" v-model="campaign.ImagesDescription[index]" placeholder="Введите текст">-->
        <!--              </b-form-input>-->
        <!--            </b-form-group>-->
        <!--          </div>-->
        <!--        </div>-->
        <h2 class="app-header-bold">{{ getPostLabel() }}</h2>
        <p class="app-new-info">
          Лучше загрузить разные варианты
          <br class="display__none"/>рекламы, это повлияет <br class="display__block"/>на эффективность <br
          class="display__none"/>рекламы. <a href="#" class="app-new-action-text">Подробнее</a></p>

        <b-form-group class="app-new-creative-uploads" :label-cols="label_cols" :state="validateImagesSmall()"
                      id="input-group-main" label-for="input-horizontal">

          <div id="image-small-block">
            <div v-for="(Img, index) in campaign.ImagesSmall" :key="'Key' + index" :style="imageSmallStyle()">
              <div class="post-image-wrapper">
                <div class="icon-div-image">
                  <b-icon @click.stop="removeImageSmall(Img , index)" class="x-button" icon="x"></b-icon>
                </div>
                <div class="icon-div-number">{{ index + 1 }}</div>
                <img id="preview-small" :src="campaign.ImagesSmallSrc[index]" alt="img"/>
              </div>
            </div>
            <div v-if="campaign.ImagesSmall.length < 5" key="s0" :style="imageSmallStyle()">
              <input style="display: none" type="file" multiple accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                     @change="onSmallFileSelected" ref="smallFileInput">
              <div @click="$refs.smallFileInput.click()" id="load-frame-small">
                <h5 id="load-file">Загрузить<br>фото или видео</h5>
                <p id="file-size">Размер<br>1080х1080рх</p>
              </div>
            </div>
          </div>
          <b-form-invalid-feedback class="error-message">
            Необходим минимум один файл
          </b-form-invalid-feedback>
        </b-form-group>
        <!--        <div v-if="isCreative()" class="app-new-creative-uploads">-->
        <!--          <div v-for="(Image, index) in ImagesSmall" :key="Image.name">-->
        <!--            <b-form-group :label="textOnImage(index)" :label-cols="label_cols" :content-cols="content_cols"-->
        <!--              id="input-group-main" label-for="input-horizontal">-->
        <!--              <b-form-input class="form-input" v-model="campaign.ImagesSmallDescription[index]"-->
        <!--                placeholder="Введите текст"></b-form-input>-->
        <!--            </b-form-group>-->
        <!--          </div>-->
        <!--        </div>-->
      </div>
      <b-form-group :label="'Описание для ' + ImagesSmallDescription[index] + ' варианта'" :label-cols="label_cols"
                    :content-cols="content_cols" id="input-group-main" label-for="input-horizontal"
                    style="margin-top: -14px"
                    v-for="(Img , index) in campaign.ImagesSmall" :key="Img.name + index"
      >
        <b-form-textarea class="form-input" style="height: 132px; padding-top: 6px"
                         v-model="campaign.ImagesSmallDescription[index]"
                         placeholder="Введите текст"></b-form-textarea>
      </b-form-group>
      <div v-if="campaign.ImagesSmall.length > 0">
        <h2>Предпросмотр</h2>
        <div class="creative-message" @click="$modal.show('stories-modal')">
          <b-row>
            <b-col cols="10">
              Как будет выглядеть реклама
            </b-col>
            <b-col cols="2" class="right-pointer">
              <svg width="7" height="12" viewBox="0 0 7 12" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M1 1L5.5 6L1 11" stroke="black" stroke-width="2" stroke-linecap="round"/>
              </svg>
            </b-col>
          </b-row>

        </div>
      </div>

      <div class="bottom__block">
        <div style="max-width: 145px; margin-right: 20px">
          <b-button type="button"
                    class="app-new-submit-button"
                    :class="{'disabled-look': validateImagesSmall() && validateImages()}"
                    @click="sendData">
            {{ isEdit ? "Назад" : "Продолжить" }}
          </b-button>
        </div>
        <div class="app-new-small-text-fit" style="max-width: 160px">
          <p class="text-muted text-pt-desktop" style="margin-top: 3px">Заполните все данные, <br/> чтобы продолжить</p>
        </div>
      </div>

    </b-form>
  </div>
</template>

<script>

import store from '@src/store/store'
import {APP_UI_URL} from '@src/constants'

export default {
  name: 'Step2',
  props: ['label_cols', 'content_cols', 'campaign', 'isEdit', 'windowInnerWidth'],
  data: function () {
    return {
      store, // fixme
      imageNames: [],
      ImagesSmall: [],
      Images: [],
      ImagesSmallDescription: ['первого', 'второго', 'третьего', 'четвёртого', 'пятого']
    }
  },
  methods: {
    sendData() {
      // if (this.validateImagesSmall() && this.validateImages()) {
      //   return
      // }
      this.$router.push('/campaign-step-3')
      // this.$emit('next', {
      //   imageNames: this.imageNames,
      //   ImagesSmall: this.ImagesSmall,
      //   Images: this.Images
      // })
    },
    showModalPopup() {
      this.$modal.show('popup-show')
    },
    imageBlockStyle() {
      if (this.windowInnerWidth <= 600) {
        const width = (this.windowInnerWidth - 68) / 2
        return {
          width: width + 'px',
          height: width * 1.75 + 'px'
        }
      }
    },
    imageSmallStyle() {
      if (this.windowInnerWidth <= 600) {
        const width = (this.windowInnerWidth - 68) / 2
        return {
          width: width + 'px',
          height: width + 'px'
        }
      }
    },
    getImageByName(name) {
      const uID = this.campaign.UserId
      const cID = this.campaign.Id
      return `${APP_UI_URL}images/${uID}/${cID}${name}`
    },
    removeImageSmall(Image, index) {
      this.campaign.ImagesSmall.splice(index , 1)
      this.campaign.ImagesSmallDescription.splice(index, 1)
      this.campaign.ImagesSmallSrc = []
      for (let i = 0; i < this.campaign.ImagesSmall.length; i++) {
        const reader = new FileReader()
        reader.onload = () => {
          this.campaign.ImagesSmallSrc.push(reader.result)
        }
        reader.readAsDataURL(this.campaign.ImagesSmall[i])
      }
    },
    getPostLabel() {
      if (!this.isCreative()) {
        return 'Для поста в ленте'
      }
      return 'Для поста в ленте'
    },
    getStoriesLabel() {
      if (!this.isCreative()) {
        return 'Для сториз'
      }
      return 'Для сториз'
    },

    isCreative() {
      return this.campaign.CreativeStatus === 'Создать рекламные креативы'
    },
    validateImagesSmall() {
      return this.ImagesSmall.length === 0
    },
    validateImages() {
      return this.Images.length === 0
    },
    onFileSelected(e) {
      const selectedFiles = e.target.files
      const len = Math.min(selectedFiles.length, 5 - this.campaign.Images.length)
      for (let i = 0; i < len; i++) {
        this.campaign.Images.push(selectedFiles[i])
      }
      this.campaign.ImagesSrc = []
      for (let i = 0; i < this.campaign.Images.length; i++) {
        const reader = new FileReader()
        reader.onload = () => {
          this.campaign.ImagesSrc.push(reader.result)
        }
        reader.readAsDataURL(this.campaign.Images[i])
      }
      e.target.value = ''
    },
    onSmallFileSelected(e) {
      const selectedFiles = e.target.files
      const len = Math.min(selectedFiles.length, 5 - this.ImagesSmall.length)
      for (let i = 0; i < len; i++) {
        this.campaign.ImagesSmall.push(selectedFiles[i])
      }
      this.campaign.ImagesSmallSrc = []
      for (let i = 0; i < this.campaign.ImagesSmall.length; i++) {
        const reader = new FileReader()
        reader.onload = () => {
          this.campaign.ImagesSmallSrc.push(reader.result)
        }
        reader.readAsDataURL(this.campaign.ImagesSmall[i])
      }
      e.target.value = ''
    },
    remove(arr, img) {
      for (let i = 0; i < arr.length; i++) {
        if (arr[i].name === img.name) {
          arr.splice(i, 1)
        }
      }
    },
    removeImage(Image , index) {
      this.campaign.Images.splice(index , 1)
      this.campaign.ImagesSrc = []
      for (let i = 0; i < this.campaign.Images.length; i++) {
        const reader = new FileReader()
        reader.onload = () => {
          this.campaign.ImagesSrc.push(reader.result)
        }

        reader.readAsDataURL(this.campaign.Images[i])
      }
    },
    textOnSlide(index) {
      return `Текст на слайде ${index + 1}`
    },
    textOnImage(index) {
      return `Текст на картинке ${index + 1}`
    }
  }
}
</script>

<style lang="scss">
@import '@src/assets/styles/vars.scss';

.app-new-creative-uploads {
  label[for=input-horizontal] {
    display: none;
  }
}

//#input-group-main{
//  margin-bottom: 40px !important;
//}
#image-small-block {
  width: 100%;
  display: grid;
  grid-template-columns: repeat(3, minmax(160px, 180px));
  justify-content: space-between;
  align-items: center;
  grid-gap: 20px;

  & > div {
    margin-top: 8px;
    width: 180px;
    height: 180px;
  }
}

#preview {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 12px;
}

#preview-small {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 20px;
}

#image-block {
  width: 100%;
  display: grid;
  grid-template-columns: repeat(3, minmax(160px, 180px));
  justify-content: space-between;
  align-items: center;
  grid-gap: 20px;
  //overflow-x: scroll;
  //margin-top: -30px;
  & > div {
    width: 180px;
    //width: 100%;
    height: 316px;
    margin-top: 8px;
    position: relative;
  }
}

#load-file {
  font-style: normal;
  font-weight: bold;
  font-size: 14px;
  line-height: 18px;
  margin: 27px auto 18px;
  text-align: center;
  color: $black;
}

#file-size-big {
  font-style: normal;
  font-weight: normal;
  font-size: 12px;
  line-height: 20px;
  /* bottom: 0;//TODO почему то не работает */
  margin: 0;
  position: absolute;
  bottom: 28px;
  width: 100%;
  text-align: center;
  color: $gray;
}

#file-size {
  font-style: normal;
  font-weight: normal;
  font-size: 12px;
  line-height: 20px;
  bottom: 28px;
  text-align: center;
  color: #767676;
  position: absolute;
  margin: 0;
  width: 100%;
}

.stories-image-wrapper {
  width: 100%;
  height: 100%;
  position: relative;
}

.icon-div-image {
  position: absolute;
  top: -16px;
  right: -16px;
  width: 40px;
  height: 40px;
  border-radius: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f3f3f3;
  cursor: pointer;
}

@media (max-width: 600px) {
  #image-block {
    grid-template-columns: repeat(2, minmax(120px, 300px));

    & > div {
      position: relative;
    }
  }
  #load-file {
    margin: 25px auto 16px;
  }
  #file-size {
    bottom: 21px;
  }
  #file-size-big {
    //margin: 0;
    //position: absolute;
    //bottom: 28px;
    //width: 100%;
    //text-align: center;
  }

  #image-small-block {
    grid-template-columns: repeat(2, minmax(120px, 300px));

    & > div {

    }
  }
}

#load-frame {
  border: 2px dashed #CCCCCC;
  border-radius: 16px;
  width: 100%;
  height: 100%;
}

#load-frame-small {
  border: 2px dashed #CCCCCC;
  border-radius: 16px;
  width: 100%;
  height: 100%;
  position: relative;
}

.post-image-wrapper {
  position: relative;
  width: 100%;
  height: 100%;
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
  border-radius: 16px;
  color: $light;
  width: 32px;
  height: 32px;
  font-size: 14px;
  line-height: 18px;
  margin-bottom: 0;
  left: 10px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.post-image-wrapper {
  //margin-top: -30px;

  & .icon-div-number {
    //bottom: -150px;
  }
}

</style>
