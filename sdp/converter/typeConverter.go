package converter

import (
    "github.com/jjosephy/go/sdp/contract/v1"
    "github.com/jjosephy/go/sdp/model"
)

func ConvertDeployModelToContractV1 (m model.DeployModel) (c contract.DeployContractV1) {
    // TODO: validate input
    return contract.DeployContractV1 {
        Id: "id",
        Content: "content",
    }
}

func ConvertDeployContractToModelV1 (c contract.DeployContractV1) (m model.DeployModel) {
    // TODO: validate input
    return model.DeployModel {
        Id: "id",
        Content: "content",
    }
}
