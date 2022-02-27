import { getActionMock } from "../../__mocks__/transaction.mock"
import { getHeaderParams, getMemoText } from "../action.helpers"

describe("getMemoText", () => {
  it("should return memo if it exists", () => {
    const action = getActionMock({ data: { memo: "test" } })
    expect(getMemoText(action)).toEqual("test")
  })
})

describe("getHeaderParams", () => {
  it("return default if account/name not recognized", () => {
    const unknownAccount = "newAccount"
    const receiver = "newAccount"
    expect(getHeaderParams(unknownAccount, receiver)).toEqual({
      color: "traceAccountGenericBackground",
      text: unknownAccount,
      hoverTitle: unknownAccount
    })
  })

  it("return system header params with eosio account", () => {
    expect(getHeaderParams("zswhq", "zswhq")).toEqual({
      color: "#002343",
      text: "Sy",
      hoverTitle: "zswhq"
    })
  })

  it("return forum header params with eosio account", () => {
    expect(getHeaderParams("zswhq.forum", "zswhq.forum")).toEqual({
      color: "#5449ba",
      text: "Fo",
      hoverTitle: "zswhq.forum"
    })
  })

  it("return token header params with eosio account", () => {
    expect(getHeaderParams("zswhq.token", "zswhq.token")).toEqual({
      color: "#5449ba",
      text: "Tk",
      hoverTitle: "zswhq.token"
    })
  })
})
