import {BaseApi, MethodType} from "../base_api.ts";
import {BaseResponse} from "../model/base_response.ts";
import {CustomPageDTO} from "../model/dto/custom_page_dto.ts";
import {DeviceEntity} from "../model/entity/device_entity.ts";
import {DeviceControlLightDTO} from "../model/dto/device_control_light_dto.ts";
import {DeviceRegisterDTO} from "../model/dto/device_register_dto.ts";

/**
 * # 在线设备列表
 * 获取在线设备列表，用于展示在线设备列表；
 *
 * @param paramData CustomPageDTO 分页数据
 * @constructor
 */
const GetOnlineDeviceAPI = async (paramData: CustomPageDTO): Promise<BaseResponse<DeviceEntity[]> | undefined> => {
    return BaseApi<DeviceEntity[]>(
        MethodType.GET,
        "/api/v1/device/list/online",
        null,
        paramData,
        null,
        null,
    )
}

/**
 * # 离线设备列表
 * 获取离线设备列表，用于展示离线设备列表；
 *
 * @param paramData CustomPageDTO 分页数据
 * @constructor
 */
const GetOfflineDeviceAPI = async (paramData: CustomPageDTO): Promise<BaseResponse<DeviceEntity[]> | undefined> => {
    return BaseApi<DeviceEntity[]>(
        MethodType.GET,
        "/api/v1/device/list/offline",
        null,
        paramData,
        null,
        null,
    )
}

/**
 * # 设备列表
 * 获取设备列表，用于展示设备列表；
 *
 * @param paramData CustomPageDTO 分页数据
 * @constructor
 */
const GetDeviceAPI = async (paramData: CustomPageDTO): Promise<BaseResponse<DeviceEntity[]> | undefined> => {
    return BaseApi<DeviceEntity[]>(
        MethodType.GET,
        "/api/v1/device/list",
        null,
        paramData,
        null,
        null,
    )
}

/**
 * # 未注册设备列表
 * 获取未注册设备列表，用于展示未注册设备列表；
 *
 * @param paramData CustomPageDTO 分页数据
 * @constructor
 */
const GetNoRegisterDeviceAPI = async (paramData: CustomPageDTO): Promise<BaseResponse<DeviceEntity[]> | undefined> => {
    return BaseApi<DeviceEntity[]>(
        MethodType.GET,
        "/api/v1/device/list/no-register",
        null,
        paramData,
        null,
        null,
    )
}

/**
 * # 设备控制 - 开关灯
 * 控制设备开关灯，用于控制设备开关灯；
 *
 * @constructor
 */
const DeviceLightControlAPI = async (bodyData: DeviceControlLightDTO): Promise<BaseResponse<void> | undefined> => {
    return BaseApi<void>(
        MethodType.POST,
        "/api/v1/device/control/light",
        bodyData,
        null,
        null,
        null,
    )
}

/**
 * # 设备注册
 * 设备注册，用于设备注册；
 *
 * @constructor
 */
const DeviceRegisterAPI = async (bodyData: DeviceRegisterDTO): Promise<BaseResponse<void> | undefined> => {
    return BaseApi<void>(
        MethodType.POST,
        "/api/v1/device/register",
        bodyData,
        null,
        null,
        null,
    )
}

export {
    GetOnlineDeviceAPI,
    GetOfflineDeviceAPI,
    GetDeviceAPI,
    GetNoRegisterDeviceAPI,
    DeviceLightControlAPI,
    DeviceRegisterAPI
}
