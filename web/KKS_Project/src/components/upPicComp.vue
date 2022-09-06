<template>
  <div :class="'display-area'">
    <el-row :gutter="20">
      <el-col v-for="(file, index) in fileList" :key="file" :span="4">
        <el-card :body-style="{ padding: '0px' }" shadow="hover">
          <div style="position: relative">
            <div :class="'image_show'">
              <el-image
                :src="file.url"
                fit="scale-down"
                :preview-src-list="fileUrlList"
                :initial-index="index"
              />
            </div>
            <div :class="'detailShow'">
              <p style="margin-top: 0px; margin-bottom: 0px">
                上传：{{ file.form.Uploader }} <br />
                标题：{{ file.form.Title }} <br />
                Msg: {{ file.form.Message }} <br />
              </p>

              <TagComp
                :Tags="file.form.Tags"
                v-if="file.form.Tags.length > 0"
              />
              <td v-if="file.form.Tags.length <= 0">tags:</td>

              <!-- detail上面的操作按钮 -->
              <span
                :class="
                  beginUpload == false
                    ? 'detail_button_action'
                    : 'detail_button_disabled'
                "
              >
                <el-button
                  type="info"
                  :icon="EditPen"
                  circle
                  @click="EditDetail(index)"
                  :class="'hide_button'"
                />
                <el-button
                  type="info"
                  :icon="Remove"
                  circle
                  @click="DeleteFile(index)"
                  :class="'hide_button'"
                />
              </span>
            </div>
            <!-- 上传时显示的进度条界面 -->
            <span :class="beginUpload == true ? 'progress_on' : 'progress_off'">
              <!-- <span :class="'progress_on'"> -->
              <el-progress
                type="circle"
                :percentage="file.percentage"
                :status="file.uploadStatus"
                style="opacity: 1"
              >
                <template
                  :class="
                    file.uploadStatus == 'ready'
                      ? 'progress_off'
                      : 'progressWarning_on'
                  "
                >
                  <el-button
                    type="success"
                    :icon="Check"
                    circle
                    v-if="file.uploadStatus == 'success'"
                  />
                  <el-button
                    type="warning"
                    :icon="Warning"
                    circle
                    v-if="file.uploadStatus == 'exception'"
                  />
                  <el-button
                    type="danger"
                    :icon="Close"
                    circle
                    v-if="file.uploadStatus == 'warning'"
                  />
                  <span :v-if="file.TextBox != ''"
                    ><br />{{ file.TextBox }}</span
                  >
                </template>
              </el-progress>
            </span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>

  <el-button-group>
    <el-button
      type="primary"
      :icon="Plus"
      @click="dialogFormVisible = true"
      :disabled="beginUpload"
    >
      添加图片
    </el-button>
    <el-button
      type="primary"
      @click="uploadButtonClick"
      :disabled="beginUpload"
    >
      上传所有图片
      <el-icon class="el-icon--right">
        <Upload />
      </el-icon>
    </el-button>
  </el-button-group>

  <el-dialog
    :before-close="closeButtonClick"
    :close-on-click-modal="false"
    v-model="dialogFormVisible"
    title="图片信息设置"
    width="44%"
  >
    <!-- <el-upload
			v-model:file-list="tmpfileList"
			action="none"
			list-type="picture-card"
			:on-preview="handlePictureCardPreview"
			:on-remove="handleRemove"
			:auto-upload="false"
			:on-change="onChangeTest"
			:multiple="true"
			ref="upload"
			accept=".jpg,.jpeg,.png,.gif"
			
			> -->
    <el-upload
      v-model:file-list="tmpfileList"
      action="none"
      list-type="picture-card"
      :on-preview="handlePictureCardPreview"
      :auto-upload="false"
      :multiple="true"
      ref="upload"
      accept=".jpg,.jpeg,.png,.gif"
    >
      <el-icon><Plus /></el-icon>
    </el-upload>

    <el-form :model="form" label-width="140px">
      <div class="el-upload__tip">
        下方属性会赋值给所有图片，如需添加不同属性需要多次填写表单，图片大小&lt;=16M
      </div>
      <el-form-item label="上传者" required>
        <el-input v-model="form.Uploader" placeholder="人人都可以是Karapo" />
      </el-form-item>
      <el-form-item label="标题">
        <el-input v-model="form.Title" />
      </el-form-item>
      <el-form-item label="留言">
        <el-input v-model="form.Message" />
      </el-form-item>
      <el-form-item label="标签">
        <el-input v-model="form.Tags" placeholder="以空格隔开" />
      </el-form-item>

      <el-button type="danger" :icon="Close" @click="closeButtonClick" circle />
      <el-button type="success" :icon="Check" @click="okButtonClick" circle />
    </el-form>
  </el-dialog>

  <el-dialog v-model="dialogVisible" :class="'dialogView'">
    <el-image :src="dialogImageUrl" :fit="'contain'" alt="Preview Image" />
  </el-dialog>

  <!-- 设置单个图片的表单 -->
  <el-dialog
    :before-close="EditDetailClose"
    :close-on-click-modal="false"
    v-model="PicFormVisible"
    title="图片信息设置"
    width="44%"
  >
    <el-image :src="fileList[PicIndexSelected].url" fit="scale-down" />

    <el-form :model="form" label-width="140px">
      <div class="el-upload__tip">在下方修改表单数据，标签以空格隔开</div>
      <el-form-item label="上传者" required>
        <el-input v-model="form.Uploader" />
      </el-form-item>
      <el-form-item label="标题">
        <el-input v-model="form.Title" />
      </el-form-item>
      <el-form-item label="留言">
        <el-input v-model="form.Message" />
      </el-form-item>
      <el-form-item label="标签">
        <el-input v-model="form.Tags" placeholder="以空格隔开" />
      </el-form-item>

      <el-button type="danger" :icon="Close" @click="EditDetailClose" circle />
      <el-button type="success" :icon="Check" @click="EditDetailOK" circle />
    </el-form>
  </el-dialog>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import {
  Close,
  Check,
  Plus,
  EditPen,
  Delete,
  Position,
  Warning,
  Remove,
  Upload,
} from "@element-plus/icons-vue";
import { ElMessageBox, UploadProps, UploadUserFile } from "element-plus";
import { ElMessage } from "element-plus";
import TagComp from "./TagComp.vue";
import axios from "axios";
import api from "./api_config";

const dialogFormVisible = ref(false);
const dialogImageUrl = ref("");
const dialogVisible = ref(false);

const PicFormVisible = ref(false);
const PicIndexSelected = ref(0);

const tmpfileList = ref<UploadUserFile[]>([]);
const fileList = ref<any[]>([]);
const fileUrlList = ref<string[]>([]);

const upload = ref(null);

// 在开始上传时将所有按钮禁用
const beginUpload = ref(false);

const form = reactive({
  Uploader: "",
  Title: "",
  Message: "",
  Tags: "",
});

// const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
//   console.log(uploadFile, uploadFiles)
// }

const handlePictureCardPreview: UploadProps["onPreview"] = (uploadFile) => {
  // console.log("preview")
  dialogImageUrl.value = uploadFile.url!;
  dialogVisible.value = true;
};

// const onChangeTest: UploadProps['onChange'] = (uploadFile, uploadFiles) => {
// 	console.log("onchange")
// 	console.log(uploadFile, uploadFiles)
// }

const createUpPic = (
  file,
  formData: { Uploader: string; Title: string; Message: string; Tags: string }
) => {
  var pic = {
    url: file.url,
    form: {
      Uploader: formData.Uploader,
      Title: formData.Title,
      Message: formData.Message,
      Data: file,
      FileName: file.name,
      Tags: Array<string>(),
    },
    uploadStatus: "ready",
    percentage: 0,
    TextBox: "",
  };
  if (formData.Tags != "") {
    var tmp = formData.Tags.split(" ");
    if (tmp.length != 0) {
      var i = 0;
      for (i = 0; i < tmp.length; i++) {
        if (tmp[i] != "") {
          pic.form.Tags.push(tmp[i]);
        }
      }
    }
  }
  fileUrlList.value.push(pic.url);
  return pic;
};

const okButtonClick = () => {
  // console.log(tmpfileList.value.length)
  // console.log(tmpfileList)
  if (tmpfileList.value.length == 0) {
    ElMessage({
      type: "warning",
      message: `没有要上传的东西`,
    });
    return;
  }
  // TODO：這裏要加入表單校驗
  if (form.Uploader == "") {
    ElMessage({
      type: "warning",
      message: `未填写上传者`,
    });
    return;
  }

  var i = 0;
  for (i = 0; i < tmpfileList.value.length; i++) {
    var tmp = createUpPic(tmpfileList.value[i], form);
    fileList.value.push(tmp);
  }
  form.Uploader = "";
  form.Title = "";
  form.Message = "";
  form.Tags = "";

  dialogFormVisible.value = false;
  if (upload.value != null) {
    upload.value.clearFiles();
  }
};

const closeButtonClick = () => {
  dialogFormVisible.value = false;
  upload.value.clearFiles();
};

const EditDetail = (index: number) => {
  // console.log("click edit button " + index)
  PicIndexSelected.value = index;

  form.Uploader = fileList.value[index].form.Uploader;
  form.Title = fileList.value[index].form.Title;
  form.Message = fileList.value[index].form.Message;
  form.Tags = fileList.value[index].form.Tags.join(" ");

  PicFormVisible.value = true;
};

const DeleteFile = (index: number) => {
  // console.log("click delete button" + index)
  fileList.value.splice(index, 1);
  fileUrlList.value.splice(index, 1);
};

const EditDetailClose = () => {
  // console.log("click EditDetailClose")
  PicFormVisible.value = false;
  PicIndexSelected.value = 0;

  form.Uploader = "";
  form.Title = "";
  form.Message = "";
  form.Tags = "";
};

const EditDetailOK = () => {
  var index = PicIndexSelected.value;
  fileList.value[index].form.Uploader = form.Uploader;
  fileList.value[index].form.Title = form.Title;
  fileList.value[index].form.Message = form.Message;
  fileList.value[index].form.Tags = [];

  if (form.Tags != "") {
    var tmp = form.Tags.split(" ");
    if (tmp.length != 0) {
      var i = 0;
      for (i = 0; i < tmp.length; i++) {
        if (tmp[i] != "") {
          fileList.value[index].form.Tags.push(tmp[i]);
        }
      }
    }
  }

  PicFormVisible.value = false;
  PicIndexSelected.value = 0;
};

//FIXME: 由于reader 和 axios 都是异步的，需要解决异步同步问题，否则无法正常显示发送状态； post的response还没有进行校验

const readPicAsync = (File) => {
  return new Promise((resolve, reject) => {
    var reader = new FileReader();
    var filebuf = [];

    reader.onload = () => {
      resolve(reader.result);
    };
    reader.onerror = reject;

    reader.readAsArrayBuffer(File);
  });
};

const PostSinglePic = async (form) => {
  // var bodyFormData = new FormData()
  // bodyFormData.append('title', form.form.Title)
  // bodyFormData.append('file_name', form.form.FileName)
  // bodyFormData.append('data', form.form.Data)
  // bodyFormData.append('uploader', form.form.Uploader)
  // bodyFormData.append('message', form.form.Message)
  // bodyFormData.append('tags', form.form.Tags)
  var filebuf = [];

  try {
    let arrayBuffer = await readPicAsync(form.form.Data.raw);
    var array = new Uint8Array(arrayBuffer);
    for (const a of array) {
      filebuf.push(a);
    }
    // console.log("测试async filebuf:" + filebuf)
  } catch (err) {
    console.log("PostSinglePic err:" + err);
  }

  var bodyFormData = {
    title: form.form.Title,
    file_name: form.form.FileName,
    data: filebuf,
    uploader: form.form.Uploader,
    message: form.form.Message,
    tags: form.form.Tags,
  };

  // console.log(bodyFormData)
  var ret = false;

  try {
    await axios({
      method: "post",
      url: api.upload,
      data: bodyFormData,
      // TODO: maybe add "headers:" in the fulture
      timeout: 10000,
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(function (response) {
        // console.log("PostPicFormAsync" + response.data)
        // console.log(response.status)
        // console.log(response.data.data)
        if (response.data.status) {
          ret = true;
        }
      })
      .catch(function (err) {
        console.log("err catch+" + err);
      });
  } catch (err) {
    console.log("PostPicFormAsync err:" + err);
  }
  return ret;
};

const uploadButtonClick = async () => {
  beginUpload.value = true;
  var success_num = 0;
  for (var i = 0; i < fileList.value.length; i++) {
    fileList.value[i].percentage = 20;
    fileList.value[i].form.data;
    if (checkPicSizeOversize(fileList.value[i].form.Data) == true) {
      fileList.value[i].uploadStatus = "exception";
      fileList.value[i].TextBox = "文件大小大于16M";
      continue;
    }

    fileList.value[i].percentage = 40;

    try {
      var ret = await PostSinglePic(fileList.value[i]);
      if (ret == true) {
        fileList.value[i].percentage = 80;
        fileList.value[i].uploadStatus = "success";
        success_num += 1;
        // console.log("success upload pic" + i)
        fileList.value[i].percentage = 100;
      } else {
        fileList.value[i].status = "warning";
        fileList.value[i].TextBox = "图片上传失败";
        // console.log("fail upload pic" + i)
      }
    } catch (err) {
      console.log(err);
    }
  }

  ElMessageBox.confirm(
    "上传成功" +
      success_num +
      "/" +
      fileList.value.length +
      "，选择继续上传后上传成功图片将不再显示",
    "上传结果",
    {
      confirmButtonText: "继续上传",
      type: "info",
    }
  ).then(function () {
    // TODO:在上传完成后删除全部成功的图片
    var temp = ref<object[]>([]);
    for (var i = 0; i < fileList.value.length; i++) {
      if (fileList.value[i].uploadStatus != "success") {
        temp.value.push(fileList.value[i]);
      }
    }
    // console.log(temp.value)
    fileList.value = temp.value;
    beginUpload.value = false;
  });
};

const checkPicSizeOversize = (file) => {
  if (file.size / 1024 / 1024 > 16) {
    return true;
  } else {
    return false;
  }
};
</script>
	
<style scoped>
.el-select {
  width: 300px;
}
.el-input {
  width: 300px;
}
.display-area {
  padding-left: 72px;
  padding-right: 72px;
}
.image_show {
  position: relative;
}

.el-row {
  margin-bottom: 20px;
}
.el-row:last-child {
  margin-bottom: 0;
}

.dialogView {
  width: 60%;
}

.detailShow {
  padding: 14px;
  text-align: left;
  position: relative;
}

.detail_button_action:hover {
  opacity: 1;
}

.detail_button_disabled {
  display: none;
}

.detail_button_action {
  position: absolute;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  cursor: default;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  color: #fff;
  opacity: 0;
  font-size: 20px;
  background-color: var(--el-overlay-color-lighter);
  transition: opacity var(--el-transition-duration);
}

.progress_on {
  position: absolute;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  justify-content: center;
  align-items: center;
  vertical-align: middle;
  display: flex;

  background-color: white;
  opacity: 0.85;
}

.progress_off {
  display: none;
}
.progressWarning_on {
  display: inline;
}
</style>