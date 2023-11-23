<template>
  <div class="pt-12 px-4 bg-lightBlackBg min-h-screen !text-black">
    <div class="text-3xl font-sans max-w-[608px] mx-auto bg-lightBlueBg px-8 py-[60px] rounded-lg">This is create form
      <div>
        <label class="block mb-2 text-sm font-medium text-gray-900 " for="small_size">import file</label>
        <input
          class="block w-full mb-5 text-xs text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 focus:outline-none"
          id="small_size" type="file">
      </div>
      <div class="border-blackBg70 border rounded-lg py-4 px-6">
        nothing.xsl
      </div>
      <div class="flex items-center mb-4">
        <input id="default-checkbox" type="checkbox" value=""
          class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
        <label for="default-checkbox" class="ms-2 text-sm font-medium text-gray-900 dark:text-gray-300">Default
          checkbox
        </label>
      </div>
    </div>
  </div></template>
