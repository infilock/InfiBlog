<template>
  <div class="px-[34px] py-12 bg-lightBlackBg min-h-screen !text-black">
    <div class="text-3xl font-sans max-w-[608px] mx-auto bg-lightBlueBg px-8 py-[60px] rounded-lg">
      <div>
        <div class="mb-3">
          <label for="fileInput" class="mb-2 inline-block text-lg font-bold">Import file</label>
          <input
            class="relative m-0 block w-full min-w-0 flex-auto rounded border border-solid border-neutral-300 bg-clip-padding px-3 py-[0.32rem] text-base font-normal text-neutral-700 transition duration-300 ease-in-out file:-mx-3 file:-my-[0.32rem] file:overflow-hidden file:rounded-none file:border-0 file:border-solid file:border-inherit file:bg-neutral-100 file:px-3 file:py-[0.32rem] file:text-neutral-700 file:transition file:duration-150 file:ease-in-out file:[border-inline-end-width:1px] file:[margin-inline-end:0.75rem] hover:file:bg-primary hover:file:text-white focus:border-primary focus:text-neutral-700 focus:shadow-te-primary focus:outline-none"
            type="file" name="file" id="fileInput" ref="file" @change="onChange()" accept=".xsl,.xlsx" />
        </div>

      </div>
      <div v-if="files[0]" class="border-blackBg70 border rounded-lg py-4 px-6 flex justify-start items-center gap-4">
        <img class="!w-6 !h-6" src="/icons/excel.png" alt="excel">
        <div class="flex gap-2 flex-col">
          <p class="text-base font-bold">&emsp;{{ !!file.files[0].name &&
            file.files[0].name.length > 10 ? file.files[0].name.substring(0, 10) + '...'
            : file.files[0].name ? file.files[0].name : '' }}</p>
          <span class="text-sm text-blackBg70">{{ (file.files[0].size / (1024 * 1024)).toFixed(2) }} mb</span>
        </div>
      </div>
      <div class="pt-6 pb-8">
        <p class="font-bold text-base mb-2">Tag</p>
        <Multiselect v-model="value" mode="tags" placeholder="Select employees" track-by="name" label="name"
        :close-on-select="false" :searchable="true" :options="[
          { value: 'judy', name: 'Judy' },
          { value: 'jane', name: 'Jane' },
          { value: 'john', name: 'John' },
          { value: 'joe', name: 'Joe' }
        ]">
        <template v-slot:tag="{ option, handleTagRemove, disabled }">
          <div class="multiselect-tag is-user" :class="{
            'is-disabled': disabled
          }">
            <img :src="option.image">
            {{ option.name }}
            <span v-if="!disabled" class="multiselect-tag-remove" @click="handleTagRemove(option, $event)">
              <span class="multiselect-tag-remove-icon"></span>
            </span>
          </div>
        </template>
      </Multiselect>
        <p class="font-bold text-base mb-2 mt-6">Category</p>
        <Multiselect v-model="value" mode="tags" placeholder="Select employees" track-by="name" label="name"
        :close-on-select="false" :searchable="true" :options="[
          { value: 'judy', name: 'Judy' },
          { value: 'jane', name: 'Jane' },
          { value: 'john', name: 'John' },
          { value: 'joe', name: 'Joe' }
        ]">
        <template v-slot:tag="{ option, handleTagRemove, disabled }">
          <div class="multiselect-tag is-user" :class="{
            'is-disabled': disabled
          }">
            <img :src="option.image">
            {{ option.name }}
            <span v-if="!disabled" class="multiselect-tag-remove" @click="handleTagRemove(option, $event)">
              <span class="multiselect-tag-remove-icon"></span>
            </span>
          </div>
        </template>
      </Multiselect>
      </div>

      <div class="flex justify-end gap-4 max-md:flex-col-reverse">
        <button
          class="max-md:w-full bg-transparent border-blackBg70 border transition-all duration-300 hover:shadow-lg hover:bg-primary hover:text-white hover:border-primary rounded-lg text-lg py-2 px-[42px]"
          @click="cancel()">Cancel</button>
        <!-- :disabled="v$.$invalid || v$.$errors.length || loadingTwo || loading" :isLoading="loading" -->
        <button
          class="max-md:w-full bg-primary transition-all duration-300 hover:shadow-lg rounded-lg text-white text-lg py-2 px-[42px]"
          @click="submit()">Submit</button>
      </div>
    </div>
  </div>
</template>
<script setup>
import Multiselect from '@vueform/multiselect'

import { ref, reactive } from 'vue';
const value = ref(null)
const options = ref([
  'Batman',
  'Robin',
  'Joker',
])
let files = ref([]);
let file = ref(null);
let tempAttachmentName = ref(null);

const formData = reactive({
  title: "",
  department_id: "",
  message: "",
});

const onChange = () => {
  if (file.value.files[0].size / (1024 * 1024) > 15) {
    alert("your file weight is more than 15mb!");
  } else {
    files.value = [...file.value.files];
  }
};




</script>
<style src="@vueform/multiselect/themes/default.css"></style>