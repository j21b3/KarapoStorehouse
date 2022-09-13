<template>
  <div :class="'display-area'">
    <el-row :gutter="20" v-infinite-scroll="load">
      <el-col v-for="file in fileList" :key="file" :span="4">
        <el-card :body-style="{ padding: '0px' }" shadow="hover">
          <div :class="'image_show'">
            <el-image :src="file.url" fit="scale-down" />
          </div>
          <div :class="'detailShow'">
            <p style="margin-top: 0px; margin-bottom: 0px">
              上传：{{ file.form.Uploader }} <br />
              标题：{{ file.form.Title }} <br />
              Msg: {{ file.form.Message }} <br />
            </p>
            <td v-if="file.form.Tags.length <= 0">tags:</td>
            <TagComp :Tags="file.form.Tags" v-if="file.form.Tags.length > 0" />
          </div>
        </el-card>
      </el-col>
      <p v-if="loading">Loading...</p>
      <p v-if="noMore">No more</p>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "@vue/reactivity";
import { computed } from "@vue/runtime-core";
import axios from "axios";
import { analyzeMetafile } from "esbuild";
import api from "./api_config";
import TagComp from "./TagComp.vue";

//图片页数
var page = 1;
// 加载标志位
const loading = ref(false);
// 是否滚动到尽头
const noMore = ref(false);
// 组件disabled
const disabled = computed(() => loading.value || noMore.value);

const thumbnailWidth = 186;

const fileList = ref<any[]>([]);

const createPicData = async (hexid: string) => {
  var picdata = {
    url: api.raw.thumbnail + "?pid=" + hexid + "&width=" + thumbnailWidth,
    form: {},
  };
  await axios
    .get(api.raw.detail, {
      params: {
        pid: hexid,
      },
    })
    .then((response) => {
      if (response.data.status) {
        picdata.form = {
          Uploader: response.data.data.uploader,
          Title: response.data.data.title,
          Message: response.data.data.message,
          FileName: response.data.data.file_name,
          Tags: response.data.data.tags,
        };
      }
    })
    .catch((error) => {
      console.log(error);
    });
  return picdata;
};

const load = async () => {
  var data;
  // 加载图片
  await axios
    .get(api.timeline + page)
    .then(function (response) {
      if (response.data.status != true) {
        return;
      }
      if (response.data.data.length <= 0 || response.data.data.length < 24) {
        noMore.value = true;
      }
      data = response.data.data;
    })
    .catch(function (error) {
      console.log(error);
    });
  for (var i = 0; i < data.length; i++) {
    var tmp = await createPicData(data[i]);
    console.log(tmp);

    fileList.value.push(tmp);
  }
  page += 1;
};
</script>

<style>
.display-area {
  padding-left: 72px;
  padding-right: 72px;
}
.image_show {
  position: relative;
}
.detailShow {
  padding: 14px;
  text-align: left;
  position: relative;
}
</style>