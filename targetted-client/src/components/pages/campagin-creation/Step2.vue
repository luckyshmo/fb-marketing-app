<template>
  <div>
    <slot name="header"></slot>
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
        <b-form-group label="Наличие креативов" :label-cols="label_cols" :content-cols="content_cols"
          id="input-group-main" label-for="input-horizontal">
          <b-form-radio-group v-model="company.CreativeStatus" :options="[
                                  'Есть рекламные креативы',
                                  'Создать рекламные креативы'
                              ]"></b-form-radio-group>
        </b-form-group>

        <div class="creative-message">
          <div v-if="isCreative()">
            Для создания рекламных креативов загрузите картинки и напишите текст, который будет на них отображаться. <a
              href="https://docs.google.com/document/d/1gqOkpxDJ1wNt-AYlt5Q1Et1kF8NRLRYdG-dXK7WdT1k/edit?usp=sharing"
              style="color: #6C1BD2" target="_blank">Посмотрите пример,</a> как это будет выглядеть и подробную
            инструкцию.
          </div>
          <div v-if="!isCreative()">
            <a href="https://docs.google.com/document/d/1DGJQw2lw_Y6KzyPeqA6To8u9UuiYi2Kokx_yjvTxgKE/edit?usp=sharing"
              style="color: #6C1BD2" target="_blank">Воспользуйтесь советами</a> при самостоятельном создании креативов,
            чтобы увеличить эффективность рекламной кампании.
          </div>
        </div>

        <b-form-group :label="getStoriesLabel()" :label-cols="label_cols" :state="validateImages()" id="input-group-main"
          label-for="input-horizontal" description="До 5 слайдов в сториз">
          <div id="image-block">
            <div v-for="(Image, key) in Images" :key="key">
              <div>
                <div id="icon-div-image">
                  <b-icon @click.stop="removeImage(Image)" class="x-button" icon="x"></b-icon>
                </div>
                <img id="preview" :ref="'Image'" />
              </div>
            </div>

            <div v-if="Images.length < 5">
              <input style="display: none" type="file" multiple accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                @change="onFileSelected" ref="fileInput">
              <div @click="$refs.fileInput.click()" id="load-frame">
                <!-- Обводка у кнопок -->
                <!-- ПОехал размер -->
                <p id="load-file">Загрузить<br>файл</p>
                <p id="file-size-big">Размер<br>1920х1080рх</p>
                <!-- TODO FILL -->
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
              <b-form-input class="form-input" v-model="company.ImagesDescription[index]" placeholder="Введите текст">
              </b-form-input>
            </b-form-group>
          </div>
        </div>
        <b-form-group :label="getPostLabel()" :label-cols="label_cols" :state="validateImagesSmall()"
          id="input-group-main" label-for="input-horizontal" description="До 5 слайдов в посте">

          <div id="image-block">
            <div v-for="(Image, key) in ImagesSmall" :key="key" style="width: 160px; height: 160px;">
              <div id="icon-div-image">
                <b-icon @click.stop="removeImageSmall(Image)" class="x-button" icon="x"></b-icon>
              </div>
              <img id="preview-small" :ref="'ImageSmall'" />
            </div>
            <div v-if="ImagesSmall.length < 5">
              <input style="display: none" type="file" multiple accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                @change="onSmallFileSelected" ref="smallFileInput">
              <div @click="$refs.smallFileInput.click()" id="load-frame-small">
                <p id="load-file">Загрузить<br>файл</p>
                <p id="file-size">Размер<br>1080х1080рх</p>
              </div>
            </div>
          </div>
          <b-form-invalid-feedback class="error-message">
            Необходим минимум один файл
          </b-form-invalid-feedback>
        </b-form-group>
        <div v-if="isCreative()">
          <div v-for="(Image, index) in ImagesSmall" :key="Image.name">
            <b-form-group :label="textOnImage(index)" :label-cols="label_cols" :content-cols="content_cols"
              id="input-group-main" label-for="input-horizontal">
              <b-form-input class="form-input" v-model="company.ImagesSmallDescription[index]"
                placeholder="Введите текст"></b-form-input>
            </b-form-group>
          </div>
        </div>
      </div>
      <b-form-group v-if="ImagesSmall.length > 0" label="Описание под постом в ленте" :label-cols="label_cols"
        :content-cols="content_cols" id="input-group-main" label-for="input-horizontal">
        <b-form-textarea class="form-input" style="height: 100px" v-model="company.PostDescription"
          placeholder="Введите текст"></b-form-textarea>
      </b-form-group>
      <b-button class="app-new-submit-button" type="submit">
        {{isEdit ? "Назад":"Продолжить"}}
      </b-button>
    </b-form>
  </div>
</template>

<script>
import store from '@/../store/store'
import { required, maxLength, minLength } from 'vuelidate/lib/validators'
import {APP_UI_URL} from '@/constants'

export default {
    name: 'Step2',
    props: ['label_cols', 'content_cols', 'company', 'isEdit'],
    data: function () {
        return {
          store, //fixme
          imageNames: [],
          ImagesSmall: [],
          Images: [],
        }
      },
      methods: {
        getImageByName(name) {
          const uID = this.company.UserId
          const cID = this.company.Id
          return `${APP_UI_URL}images/${uID}/${cID}${name}`
        },
        removeImageSmall(Image) {
          this.remove(this.ImagesSmall, Image)
          for (let i = 0; i < this.ImagesSmall.length; i++) {
            const reader = new FileReader()
            reader.onload = () => {
              this.$refs.ImageSmall[i].src = reader.result
            }

            reader.readAsDataURL(this.ImagesSmall[i])
          }
        },
        getPostLabel() {
          if (!this.isCreative()) {
            return 'Креативы для поста в ленте'
          }
          return 'Картинки для поста в ленте'
        },
        getStoriesLabel() {
          if (!this.isCreative()) {
            return 'Креативы для Сториз'
          }
          return 'Картинки для Сториз'
        },

        isCreative() {
          return this.company.CreativeStatus === 'Создать рекламные креативы'
        },
        validateImagesSmall() {
          return this.ImagesSmall.length === 0
        },
        validateImages() {
          return this.Images.length === 0
        },
        onFileSelected(e) {
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
        onSmallFileSelected(e) {
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
        remove(arr, img) {
          for (let i = 0; i < arr.length; i++) {
            if (arr[i].name === img.name) {
              arr.splice(i, 1)
            }
          }
        },
        removeImage(Image) {
          this.remove(this.Images, Image)
          for (let i = 0; i < this.Images.length; i++) {
            const reader = new FileReader()
            reader.onload = () => {
              this.$refs.Image[i].src = reader.result
            }

            reader.readAsDataURL(this.Images[i])
          }
        },
        textOnSlide(index) {
          return `Текст на слайде ${index + 1}`
        },
        textOnImage(index) {
          return `Текст на картинке ${index + 1}`
        },
      }
}
</script>

<style>

</style>