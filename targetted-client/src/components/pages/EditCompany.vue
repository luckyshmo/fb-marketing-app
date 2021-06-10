<template>
    <div>
        <b-form @submit.prevent="createCompany()">
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
                v-model="form.CompanyName"
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
                    v-model="form.CompnayPurpose"
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
                v-model="form.CompanyField"
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
                v-model="form.BusinessAddress"
                placeholder="Точный адрес"
                ></b-form-input>
            </b-form-group>
        </b-form>     
    </div>
</template>
<script>
import store from '../../../store/store'
export default {
    name: 'EditCompany',
    props: {
      form: {
        type: Object
      },
    },
    data() {
        return {
            store,
            label_cols: this.getWidth().label,
            content_cols: this.getWidth().content,
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
    }
}
</script>
<style>
    
</style>