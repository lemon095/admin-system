<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="兑换码" prop="redeemCode">
                        <el-input
                            v-model="queryParams.redeemCode"
                            placeholder="请输入兑换码"
                            clearable
                            size="mini"
                            style="width: 180px;"
                        />
                    </el-form-item>
                    <el-form-item label="礼包名称" prop="giftName">
                        <el-input
                            v-model="queryParams.giftName"
                            placeholder="请输入礼包名称"
                            clearable
                            size="mini"
                            style="width: 180px;"
                        />
                    </el-form-item>
                    <el-form-item label="状态" prop="isEnable">
                        <el-select
                            v-model="queryParams.isEnable"
                            placeholder="请选择状态"
                            clearable
                            size="mini"
                            style="width: 120px;"
                        >
                            <el-option label="启用" :value="1" />
                            <el-option label="禁用" :value="0" />
                        </el-select>
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:giftPack:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:giftPack:edit']"
                                type="success"
                                icon="el-icon-edit"
                                size="mini"
                                :disabled="single"
                                @click="handleUpdate"
                        >修改
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:giftPack:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="giftPackList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                    <el-table-column label="ID" align="center" prop="id"/>
                    <el-table-column label="兑换码" align="center" prop="redeemCode" :show-overflow-tooltip="true"/>
                    <el-table-column label="礼包名称" align="center" prop="giftName"
                                                 :show-overflow-tooltip="true"/>
                            <el-table-column sortable align="left" label="创建时间" prop="createdAt" width="180">
                                <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
                            </el-table-column>
                                                <el-table-column label="金币数量" align="center" prop="goldCoinNum"
                                                 :show-overflow-tooltip="true"/><el-table-column label="银币数量" align="center" prop="silverCoinNum"
                                                 :show-overflow-tooltip="true"/>
                                                 <!-- <el-table-column label="道具" align="center" prop="item":show-overflow-tooltip="true"/> -->
                                <el-table-column label="有效期" align="center" prop="startAt"
                                                 :show-overflow-tooltip="true">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.startAt) }}-{{ parseTime(scope.row.endAt) }}</span>
                                    </template>
                                </el-table-column>
                                <el-table-column label="初始数量" align="center" prop="giftNum"/>
                                <el-table-column label="剩余数量" align="center">
                                    <template #default="scope">
                                        {{ scope.row.giftNum - scope.row.redeemedNum }}
                                    </template>
                                </el-table-column>
                                <el-table-column label="详情" align="center" prop="giftDesc" :show-overflow-tooltip="true">
                                    <template #default="scope">
                                        <el-button @click="detail(scope.row.id)">详情</el-button>
                                    </template>
                                </el-table-column>
                            <el-table-column label="状态" align="center" prop="isEnable" :show-overflow-tooltip="true">
                                <template #default="scope">
                                    <el-button v-if="scope.row.isEnable === 0" type="success" @click="giftPackStatus(scope.row.id)">启用</el-button>
                                    <el-button v-else type="danger" @click="giftPackStatus(scope.row.id)">禁用</el-button>
                                </template>
                            </el-table-column>
                            <el-table-column
                                        label="创建人"
                                        align="center"
                                        prop="created_by"
                                        :show-overflow-tooltip="true"
                                      />
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                          <!-- <el-button
                             slot="reference"
                             v-permisaction="['admin:giftPack:edit']"
                             size="mini"
                             type="text"
                             icon="el-icon-edit"
                             @click="handleUpdate(scope.row)"
                           >详情
                           </el-button> -->
                         <el-popconfirm
                            class="delete-popconfirm"
                            title="确认要删除吗?"
                            confirm-button-text="删除"
                            @confirm="handleDelete(scope.row)"
                         >
                            <el-button
                              slot="reference"
                              v-permisaction="['admin:giftPack:remove']"
                              size="mini"
                              type="text"
                              icon="el-icon-delete"
                            >删除
                            </el-button>
                         </el-popconfirm>
                        </template>
                    </el-table-column>
                </el-table>

                <pagination
                        v-show="total>0"
                        :total="total"
                        :page.sync="queryParams.pageIndex"
                        :limit.sync="queryParams.pageSize"
                        @pagination="getList"
                />

                <!-- 添加或修改对话框 -->
                <el-dialog :title="title" :visible.sync="open" width="800px">
                    <el-form ref="form" :model="form" :rules="rules" :disabled="form.mode === 'view'" label-width="80px">
                        <h3>创建礼包兑换码</h3>
                                    <el-form-item label="礼包名称" prop="giftName">
                                        <el-input v-model="form.giftName" placeholder="礼包名称" style="width: 30%;"/>
                                    </el-form-item>
                                    <el-form-item label="礼包描述" prop="giftDesc">
                                        <el-input v-model="form.giftDesc" placeholder="礼包描述" type="textarea" :rows="2" style="width: 40%;"/>
                                    </el-form-item>
                                    <el-form-item label="有效时间" prop="timeRange">
                                            <el-date-picker
                                                v-model="timeRange"
                                                type="datetimerange"
                                                range-separator="至"
                                                start-placeholder="开始日期"
                                                end-placeholder="结束日期"
                                                align="right">
                                            </el-date-picker>
                                    </el-form-item>
                                    <h3>兑换码配置</h3>
                                    <el-form-item prop="giftNum">
                                        <template #label><span>数量：</span></template>
                                        <el-input-number v-model.number="form.giftNum" :min="1" placeholder="礼包码配置数量" style="width: 200px; margin-right: 20px;"/>
                                        <span style="margin-right: 8px;">单用户兑换次数：</span>
                                        <el-input-number v-model.number="form.giftLimit" :min="1" placeholder="单用户兑换次数" style="width: 200px;"/>
                                    </el-form-item>
                                    <el-form-item label="兑换码" prop="redeemCode" v-show="form.mode === 'view'">
                                        <el-input v-model="form.redeemCode" placeholder="兑换码" style="width: 200px;"/>
                                    </el-form-item>
                                    <h3>礼包配置</h3>
                                    <el-form-item prop="goldCoinNum" label-width="100px">
                                        <template #label><span>金币数量：</span></template>
                                        <el-input-number v-model.number="form.goldCoinNum" placeholder="金币数量" style="width: 200px; margin-right: 20px;margin-left: 5px;"/>
                                        <span style="margin-right: 8px;">银币数量：</span>
                                        <el-input-number v-model.number="form.silverCoinNum" placeholder="银币数量" style="width: 200px;margin-right: 1px;margin-left: 5px;"/>
                                    </el-form-item>
                                    <el-form-item label="道具" prop="item">
                                        <el-form-item
                                            v-for="(item, index) in form.items"
                                            :key="index"
                                        >
                                            <!-- <el-select v-model="item.type" placeholder="请选择道具类型" style="width: 150px; margin-right: 10px;">
                                                <el-option v-for="option in typeOptions" :key="option.id" :label="option.name" :value="option.id" />
                                            </el-select> -->
                                            <el-cascader
                                                v-model="item.value"
                                                :options="typeOptions"
                                                placeholder="请选择道具类型"
                                                clearable
                                            />
                                            <span style="margin-right: 1px;margin-left: 10px;">数量：</span>
                                            <el-input-number v-model="item.num" :min="1" :max="100" style="width: 100px; margin-right: 10px;"/>
                                            <el-button type="danger" link @click="removeItem(index)" v-show="form.items.length > 1 && form.mode !== 'view'">
                                                删除
                                            </el-button>
                                        </el-form-item>
                                        <el-button type="primary" link @click="addItem" v-show="form.mode !== 'view'">添加道具</el-button>
                                    </el-form-item>
                    </el-form>
                    <div slot="footer" class="dialog-footer">
                        <el-button type="primary" @click="submitForm">确 定</el-button>
                        <el-button @click="cancel">取 消</el-button>
                    </div>
                </el-dialog>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {addGiftPack, delGiftPack, getGiftPack, listGiftPack, updateGiftPack, updateGiftPackStatus} from '@/api/admin/bunker-gift-pack'
    import { loadItemOptions } from '@/api/admin/bunker-item'
    import { formatDate } from '@/utils/index'
    import { numberValidator } from '@/utils/validate'

    export default {
        name: 'GiftPack',
        components: {
        },
        computed: {
            timeRange: {
                get() {
                return this.form.startAt && this.form.endAt
                    ? [this.form.startAt, this.form.endAt]
                    : []
                },
                set(val) {
                if (val && val.length === 2) {
                    this.form.startAt = val[0]
                    this.form.endAt = val[1]
                } else {
                    this.form.startAt = ''
                    this.form.endAt = ''
                }
                }
            }
        },
        data() {
            return {
                // 遮罩层
                loading: true,
                // 选中数组
                ids: [],
                // 非单个禁用
                single: true,
                // 非多个禁用
                multiple: true,
                // 总条数
                total: 0,
                // 弹出层标题
                title: '',
                // 是否显示弹出层
                open: false,
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                giftPackList: [],
                
                // 关系表类型
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    redeemCode: null,
                    giftName: null,
                    isEnable: null,
                },
                // 表单参数
                form: {},
                formatDate,
                // 表单校验
                rules: {
                    giftName: [{ required: true, message: '请输入名称', trigger: 'blur' }],
                    giftNum: [{ validator: numberValidator, trigger: 'blur'}],
                    giftLimit: [{ validator: numberValidator, trigger: 'blur'}],
                }
        }
        },
        created() {
            this.getList()
            this.loadItemType()
            },
        methods: {
            async loadItemType() {
                if (this.loaded) {
                    return
                }
                try {
                    const response = await loadItemOptions()
                    if (response.code === 200) {
                        this.typeOptions = response.data
                        this.loaded = true
                    } else {
                        this.msgError(response.msg)
                    }
                } catch (err) {
                    this.msgError('获取类型失败')
                }
            },
            addItem() {
                this.form.items.push({ value: [], num: 1 })
            },
            removeItem(index) {
                this.form.items.splice(index, 1)
            },
            giftPackStatus(id) {
                updateGiftPackStatus(id).then(response => {
                    if (response.code === 200) {
                        this.msgSuccess(response.msg)
                        this.getList()
                    }
                })
            },
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listGiftPack(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.giftPackList = response.data.list
                        this.total = response.data.count
                        this.loading = false
                    }
                )
            },
            // 取消按钮
            cancel() {
                this.open = false
                this.reset()
            },
            // 表单重置
            reset() {
                this.form = {
                mode: "view",
                id: undefined,
                giftName: undefined,
                giftNum: 1,
                giftLimit: 1,
                redeemCode: undefined,
                giftDesc: undefined,
                goldCoinNum: undefined,
                silverCoinNum: undefined,
                startAt: undefined,
                endAt: undefined,
                items: [{ value: [], num: 1 }]
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            // 关系
            // 文件
            /** 搜索按钮操作 */
            handleQuery() {
                this.queryParams.pageIndex = 1
                this.getList()
            },
            /** 重置按钮操作 */
            resetQuery() {
                this.dateRange = []
                this.resetForm('queryForm')
                this.handleQuery()
            },
            /** 新增按钮操作 */
            handleAdd() {
                this.reset()
                this.open = true
                this.title = '添加兑换码管理'
                this.isEdit = false
                this.form.mode = "add"
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.id)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 修改按钮操作 */
            handleUpdate(row) {
                this.reset()
                const id =
                row.id || this.ids
                getGiftPack(id).then(response => {
                    this.form = response.data
                    this.form.mode = 'edit'
                    this.open = true
                    this.title = '修改兑换码管理'
                    this.isEdit = true
                })
            },
            detail(id) {
                this.reset()
                getGiftPack(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.form.mode = 'view'
                    this.title = '兑换码详情'
                    this.form.items = JSON.parse(this.form.item)
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.mode == "view") {
                            this.open = false
                            return
                        }
                        if (this.form.id !== undefined) {
                            updateGiftPack(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            this.form.item = JSON.stringify(this.form.items)
                            addGiftPack(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        }
                    }
                })
            },
            /** 删除按钮操作 */
            handleDelete(row) {
                var Ids = (row.id && [row.id]) || this.ids

                this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(function () {
                      return delGiftPack( { 'ids': Ids })
                }).then((response) => {
                   if (response.code === 200) {
                     this.msgSuccess(response.msg)
                     this.open = false
                     this.getList()
                   } else {
                     this.msgError(response.msg)
                   }
                }).catch(function () {
                })
            }
        }
    }
</script>
