<template>
   <h2 id="h2">Доступ к странице Facebook и Instagram</h2>
    <b-form>
        <div v-if="!(store.getters.GET_FB_PAGES.length > 0) && !isRequestSent && !pageSubmitted">
            <p>Раздайте доступ к вашей Facebook и Instagram странице<br>для запуска и управления рекламой от имени ваших страниц. </p>
            <b-button
                v-if="!(store.getters.GET_FB_PAGES.length > 0)"

                class="main-button"
                @click="loginFB"
            >
                У меня есть бизнес-аккаунт
            </b-button>
            <!-- //TODO ситуация с осутсвием страниц -->
            <popup
            v-if="isInfoPopupVisible"
            popupTitle="Инструкция по созданию бизнесс-аккаунта"
            @closePopup="closeInfoPopup"
            >
                <div style="text-align: left; padding: 30px">
                    <h1 style="margin=0px !important;">Как создать<br>бизнес-аккаунт</h1>
                    <p>Профессиональный аккаунт в Instagram необходим для того,<br>чтобы мы могли запустить и управлять рекламными кампаниями <br>от имени вашего бизнеса. </p>
                    <div class="num-wrapper">
                        <div class="num">
                            <div class="num-num">
                                1
                            </div>
                        </div>
                        <p class="num-text">Откройте мобильное приложение Инстаграм.</p>
                    </div>
                    <div class="num-wrapper">
                        <div class="num">
                            <div class="num-num">
                                2
                            </div>
                        </div>
                        <p class="num-text">Зайдите в настройки профиля. </p>
                    </div>
                    <div class="num-wrapper">
                        <div class="num">
                            <div class="num-num">
                                3
                            </div>
                        </div>
                        <p class="num-text">Нажмите на вкладку “Переключиться на профессиональный аккаунт”.</p>
                    </div>
                    <div class="num-wrapper">
                        <div class="num">
                            <div class="num-num">
                                4
                            </div>
                        </div>
                        <p class="num-text">Дайте разрешение приложению Инстаграма управлять вашими Фейсбук-страницами. Не вашим личным аккаунтом, а именно страницами в Фейсбуке, к которым вы имеете доступ. Вы можете выбрать или существующую страницу в Facebook, к которой у вас есть доступ, или создать новую. </p>
                    </div>
                    <div class="num-wrapper">
                        <div class="num">
                            <div class="num-num">
                                5
                            </div>
                        </div>
                        <p class="num-text">Настройка контактной информации для профиля Инстаграма. Укажите в соответствующие поля рабочий email компании, номер телефона и физический адрес (если есть).</p>
                    </div>
                    <div class="num-wrapper">
                        <div class="num">
                            <div class="num-num">
                                6
                            </div>
                        </div>
                        <p class="num-text">Когда вы прошли все шаги и включили профессиональный аккаунт, то, пожалуйста, вернитесь на сайт client.targetted.online и нажмите кнопку “у меня есть профессиональный аккаунт” и следуйте дальнейшим инструкциям на сайте.</p>
                    </div>
                </div>
            </popup>
            <b-button
                v-if="!(store.getters.GET_FB_PAGES.length > 0)"
                class="main-button"
                id="primary-under"
                @click="showPopupInfo"
            >
                Нет бизнес-аккаунта
            </b-button>
        </div>
        <div v-if="store.getters.GET_FB_PAGES.length > 0 && !isRequestSent && !pageSubmitted">
            <div>
                <p>Выберите страницу которую хотите привязать</p>
                <!-- //TODO SELECT -->
                <b-form-group
                label="Выберите страницу"
                :label-cols="label_cols"
                :content-cols="content_cols"
                id="input-group-main"
                label-for="input-horizontal"
                >
                    <b-form-radio-group
                        v-model="company.FbPageId"
                        :options="store.getters.GET_FB_PAGES"
                    ></b-form-radio-group>
                </b-form-group>
            </div>
            <b-button
                class="main-button"
                @click="sendFbRequest()"
            >
                Привязать
            </b-button>
        </div>
        <div v-if="isRequestSent && !pageSubmitted">
            <p>Зайди в аккаунт на Facebook и подтверди привязку страницы в сообщениях</p>
            <b-button
                class="main-button"
                target="_blank"
                rel="noopener noreferrer"
                :href=getFBRedirect()
            >
            Перейти в facebook
            </b-button>
            <b-button

                class="main-button"
                id="primary-under"
                @click="checkPageSubmitted()"
            >
                Я подтвердил в сообщениях
            </b-button>
        </div>
        <div v-if="pageSubmitted">
            <div class="c-status" style="margin-top: 30px; max-width: 800px">
                <div class="elipse" id="green"></div>
                <p class="c-status-text" id="c-status-text-u" style="text-align: left;"
                >Страница {{company.FbPageId}} привязана к targetted</p>
            </div>
        </div>
        <b-button
            v-if="store.getters.GET_FB_PAGES.length > 0 ||
            (isRequestSent && pageSubmitted) ||
            (store.getters.GET_FB_PAGES.length == 0 || isRequestSent)"

            class="main-button-grey"
            style="margin-top: 30px; background: #F3F3F3; color: black"
            @click="logout"
        >
            Привязать другой аккаунт
        </b-button>
    </div>

    <h2 id="h2">Основное</h2>

    <b-form-group
        label="Название кампании"
        :label-cols="label_cols"
        :content-cols="content_cols"
        id="input-group-main"
        label-for="input-horizontal"
    >
        <b-form-input
        class="form-input"
        id="c-name"
        v-model="company.CompanyName"
        :disabled="isEdit"
        :state="validateState('CompanyName')"
        placeholder="Введите название"
        @click="resetNameErr()"
        ></b-form-input>
        <b-form-invalid-feedback
        class="error-message"
        id="c-name">
            Название должно быть между 3 и 30 символами
        </b-form-invalid-feedback>
        <small v-if="isCompanyExist" class="error-message">
            Кампания с таким иминем уже создана
        </small>
    </b-form-group>

    <b-form-group
        label="Цель кампании"
        :label-cols="label_cols"
        :content-cols="content_cols"
        id="input-group-main"
        label-for="input-horizontal"
    >
        <b-form-radio-group
            v-model="company.CompnayPurpose"
            :disabled="isEdit"
            :options="[
                'Сообщения в директ',
                'Подписки в instagram',
                'Заявки через форму обратной связи',
                'Целевое действие на сайте'
            ]"
        ></b-form-radio-group>
    </b-form-group>

    <b-form-group
        label="Сфера деятельности"
        :label-cols="label_cols"
        :content-cols="content_cols"
        id="input-group-main"
        label-for="input-horizontal"
    >
        <b-form-input
        class="form-input"
        :disabled="isEdit"
        v-model="company.CompanyField"
        :state="validateState('CompanyField')"
        placeholder="Вставьте ссылку"
        ></b-form-input>
        <b-form-invalid-feedback
        class="error-message">
            Обязательное поле
        </b-form-invalid-feedback>
    </b-form-group>

    <b-form-group
        label="Адрес бизнеса"
        :label-cols="label_cols"
        :content-cols="content_cols"
        id="input-group-main"
        label-for="input-horizontal"
        description="Нужен только для офлайн бизнеса"
    >
        <b-form-input
        class="form-input"
        :disabled="isEdit"
        v-model="company.BusinessAddress"
        placeholder="Точный адрес"
        ></b-form-input>
    </b-form-group>
    </b-form>
                    
</template>

<script>
export default {
    name: 'Step1',
    props: ['label_cols', 'content_cols'],
    data:() => ({
        label_cols: this.getWidth().label,
        content_cols: this.getWidth().content,
        company: {
            FbPageId: '',
            UserId: '',
            Id: '',
            CompanyName: '',
            CompnayPurpose: 'Сообщения в директ',
            CompanyField: '',
            BusinessAddress: '',
            ImagesDescription: [],
            ImagesSmallDescription: [],
            CreativeStatus: 'Есть рекламные креативы',
            PostDescription: '',
            DailyAmount: 0,
            Days: 0
        }
    }),
    methods: {
        validateState (name) {
            const { $dirty, $error } = this.$v.company[name]
            return $dirty ? !$error : null
        },
            resetNameErr () {
        this.isCompanyExist = false
        },
    },
    validations: {
        company: {
        CompanyName: {
            required,
            maxLength: maxLength(30),
            minLength: minLength(3)
        },
        CompanyField: {
            required
        }
        }
    },
}
</script>

<style>

</style>