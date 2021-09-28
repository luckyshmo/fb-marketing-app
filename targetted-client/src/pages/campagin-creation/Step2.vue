<template>
  <div>
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
        <b-form-group label="Наличие креативов" :label-cols="label_cols" :content-cols="content_cols"
          id="input-group-main" label-for="input-horizontal">
          <b-form-radio-group
          v-model="company.CreativeStatus"
          class="app-new-radio"
          :options="[
                                  'Есть рекламные креативы',
                                  'Создать рекламные креативы'
                              ]"></b-form-radio-group>
        </b-form-group>

        <div class="creative-message">
          Советы по самостоятельному созданию креативов
              <b-icon icon="arrow-right">huw</b-icon>
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
      <h2>{{getStoriesLabel()}}</h2>
      
      <p class="app-new-info">В одном рекламном объявлении может быть до 5 слайдов</p>

        <b-form-group :label-cols="label_cols" :state="validateImages()" id="input-group-main"
         class="app-new-creative-uploads"
          label-for="input-horizontal" description="До 5 слайдов в сториз">
          <div id="image-block">
            <div v-if="Images.length < 5" key="s0">
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

            <div v-for="(Image, key) in Images" :key="key">
              <div class="stories-image-wrapper">
                <div class="icon-div-image">
                  <b-icon @click.stop="removeImage(Image)" class="x-button" icon="x"></b-icon>
                </div>
                <div class="icon-div-number">{{key+1}}</div>
                <img id="preview" :ref="'Image'" />
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
        
        <br><br>
        <h2>{{getPostLabel()}}</h2>
        <p class="app-new-info">В одном рекламном объявлении может быть до 5 слайдов</p>

        <b-form-group class="app-new-creative-uploads" :label-cols="label_cols" :state="validateImagesSmall()"
          id="input-group-main" label-for="input-horizontal" description="До 5 слайдов в посте">

          <div id="image-block">
            <div v-if="ImagesSmall.length < 5" key="s0">
              <input style="display: none" type="file" multiple accept="Image/gif, Image/jpeg, Image/png, Image/jpg"
                @change="onSmallFileSelected" ref="smallFileInput">
              <div @click="$refs.smallFileInput.click()" id="load-frame-small">
                <p id="load-file">Загрузить<br>файл</p>
                <p id="file-size">Размер<br>1080х1080рх</p>
              </div>
            </div>

            <div v-for="(Image, key) in ImagesSmall" :key="key" style="width: 160px; height: 160px;">
              <div class="post-image-wrapper">
              <div class="icon-div-image">
                <b-icon @click.stop="removeImageSmall(Image)" class="x-button" icon="x"></b-icon>
              </div>
              <div class="icon-div-number">{{key+1}}</div>
              <img id="preview-small" :ref="'ImageSmall'" />
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


        <br><br>
        <h2>Предпросмотр</h2>
        <div class="creative-message" @click="$router.push('/preview')">
          Как будет выглядеть реклама
          <b-icon icon="arrow-right">huw</b-icon>
        </div>

       <b-row class="mt-5">
        <b-col cols="6" sm="8" md="3" lg="3" xl="3">
          <b-button  type="button"
                    class="app-new-submit-button"
                    @click="sendData">
            {{isEdit ? "Назад":"Продолжить"}}
          </b-button>
        </b-col>
        <b-col cols="6" sm="8" md="6" lg="6" xl="6" class="app-new-small-text-fit">
            <p class="text-muted">Загрузите креативы, чтобы продолжить</p>
        </b-col>
      </b-row>

    </b-form>
  </div>
</template>

<script>
import store from '@/store/store'
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
        sendData(){
            this.$emit('next', {
            imageNames: this.imageNames,
            ImagesSmall: this.ImagesSmall,
            Images: this.Images,
          })
        },
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

<style lang="scss">
  @import '@/assets/styles/vars.scss';

.app-new-creative-uploads {
  label[for=input-horizontal] {
    display: none;
  }
}
#preview{
    width: 160px;
    height: 280px;
    object-fit: cover;
    border-radius: 20px;
}
#preview-small{
    width: 160px;
    height: 160px;
    object-fit: cover;
    border-radius: 20px;
}
#image-block {
    max-width: 530px;
    display: grid;
    grid-template-columns: repeat(auto-fill,minmax(160px, 1fr));
    justify-content: space-between;
    align-items: center;
    grid-gap: 24px;
}
#load-file {
    font-style: normal;
    font-weight: 600;
    font-size: $baseFont;
    line-height: $baseLH;
    margin: 20px auto;
    text-align: center;
    color: $black;
}
#file-size-big {
    font-style: normal;
    font-weight: normal;
    font-size: 12px;
    line-height: 20px;
    /* bottom: 0;//TODO почему то не работает */
    margin-top: 140px;
    text-align: center;
    color: $gray;
}
#file-size {
    font-family: Montserrat;
    font-style: normal;
    font-weight: normal;
    font-size: 12px;
    line-height: 20px;
    bottom: 0;
    text-align: center;
    color: $gray;
}
@media (max-width: 450px) {
    #image-block {
        max-width: 200px;
    }
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
#load-frame:hover {
    background: $light;
}
#load-frame-small:hover {
    background: $light;
}
.stories-image-wrapper {
  margin-top: -30px;
}
.post-image-wrapper {
  margin-top: -30px;
}
.icon-div-image{
    position: absolute;
    margin-left: 140px;
    margin-top: 15px;
}
.icon-div-number {
  bottom: -280px;
  position: relative;
  background-color: $black;
  border-radius: 6px;
  color: $light;
  width: 34px;
  margin-bottom: 0px;
  padding: 6px 13px;
}
.post-image-wrapper {
  margin-top: -30px;

  & .icon-div-number {
    bottom: -170px;
  }
}

</style>