/*
 * --------------------------------------------------------------------------------
 * Copyright (c) 2016-NOW(至今) 筱锋
 * Author: 筱锋(https://www.x-lf.com)
 *
 * 本文件包含 dormitory-management 的源代码，该项目的所有源代码均遵循MIT开源许可证协议。
 * --------------------------------------------------------------------------------
 * 许可证声明：
 *
 * 版权所有 (c) 2016-2024 筱锋。保留所有权利。
 *
 * 有关MIT许可证的更多信息，请查看项目根目录下的LICENSE文件或访问：
 * https://opensource.org/licenses/MIT
 * --------------------------------------------------------------------------------
 * 免责声明：
 *
 * 使用本软件的风险由用户自担。作者或版权持有人在法律允许的最大范围内，
 * 对因使用本软件内容而导致的任何直接或间接的损失不承担任何责任。
 * --------------------------------------------------------------------------------
 */

/**
 * # 基本返回类型
 * 用于返回数据的基本类型，包含输出、状态码、消息、错误消息、数据；
 *
 * @template T 返回数据类型
 * @property {string} output 输出
 * @property {number} code 状态码
 * @property {string} message 消息
 * @property {string} error_message 错误消息
 * @property {T} data 数据
 */
export type BaseResponse<T> = {
    output: string;
    code: number;
    message: string;
    error_message?: string;
    data?: T;
}
