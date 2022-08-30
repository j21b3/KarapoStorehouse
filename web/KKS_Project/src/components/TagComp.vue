<template >
    <ul :class="'TagsBox'">
        <li v-for="value in Tags" :class="'Tag1'">
            <span style="color: rgb(61, 118, 153);">
                <span>
                    <!-- TODO：以后做好标签查询之后把href="url"填进去 -->
                    <a  style="color: inherit;text-decoration: none;" v-on:click="notDone">
                        #{{value}}
                    </a>
                </span>
            </span>
        </li>
    </ul>
</template>

<script>
	import { ElMessage } from 'element-plus'

export default {
    name:"TagComp",
    
    props: ['Tags'],

    props: {
        Tags: {
            type: Array,
            default:()=>[],
        }
    },


   
    methods: {

        notDone(){
            ElMessage({
                type: 'warning',
                message: `未完待续`,

            })
        },

    }
    
}
</script>

<style>
    .Tag1{
		display: inline;
		margin: 0px 12px 0px 0px;
	}

    .TagsBox{
		word-break: break-all;
		padding-inline-start: 0px;	
	}
</style>