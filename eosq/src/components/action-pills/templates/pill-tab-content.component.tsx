import { Cell } from "../../../atoms/ui-grid/ui-grid.component"
import { SearchShortcut } from "../../search-shortcut/search-shortcut"
import { AutorizationBox } from "../../authorization-box/authorization-box.component"
import { DetailLine, JsonWrapper } from "@dfuse/explorer"
import { t } from "i18next"
import { LinkStyledText, Text } from "../../../atoms/text/text.component"
import { theme, styled } from "../../../theme"

import { RamUsage } from "../../ram-usage/ram-usage.component"
import { DBOperations } from "../../db-operations/db-operations.component"
import * as React from "react"
import { DbOp, RAMOp, TableOp, Action, Authorization } from "@dfuse/client"
import { MonospaceTextLink } from "../../../atoms/text-elements/misc"
import { Links } from "../../../routes"
import { VerticalTabs } from "../../../atoms/vertical-tabs/vertical-tabs"
import { decodeDBOps } from "../../../services/dbops"
import { TraceInfo } from "../../../models/pill-templates"

const ContentWrapper: React.ComponentType<any> = styled(Cell)`
  padding: 24px 24px 24px 40px;
`

const RawWrapper = styled(Cell)`
  margin-bottom: 10px;
`

const TabContentWrapper: React.ComponentType<any> = styled(Cell)`
  overflow-y: scroll;
  max-height: 500px;

  background-repeat: repeat;
  background-size: 50px;
  background-image: url("data:image/svg+xml,%0A%3Csvg width='100px' height='100px' viewBox='0 0 100 100' version='1.1' xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink'%3E%3Cg id='Page-1' stroke='none' stroke-width='1' fill='none' fill-rule='evenodd' opacity='0.2'%3E%3Cg id='Group-4'%3E%3Crect id='Rectangle' opacity='0.368430728' x='0' y='0' width='100' height='100'%3E%3C/rect%3E%3Cpath d='M21.5851807,61.2185355 C22.0956566,60.9879704 22.3508945,60.6549319 22.3508945,59.8095266 L22.3508945,38.5719197 L21.2788952,38.5719197 L21.2788952,59.8095266 C21.2788952,60.1938018 21.1257524,60.3218935 20.7428955,60.3218935 C20.6030696,60.3335888 20.2823817,60.339945 19.8415432,60.3409619 L19.5769775,60.3409619 C19.0925701,60.339945 18.4995761,60.3335888 17.858707,60.3218935 C18.0118497,60.6549319 18.2160401,61.1416805 18.2670877,61.4234822 L18.4876475,61.4234822 C20.1100811,61.4210236 21.0231142,61.3902329 21.5851807,61.2185355 Z M5.36403547,42.3805135 L6.13908507,41.4146531 C5.20902555,40.5758795 3.42641146,39.4320974 2.00548719,38.669576 L1.25627257,39.5083495 C2.70303183,40.3471231 4.48564592,41.54174 5.36403547,42.3805135 Z M7.86602081,55.6014789 L7.86602081,41.0044394 L14.1196266,41.0044394 L14.1196266,55.5757799 L15.2049631,55.5757799 L15.2049631,39.9764789 L6.80652562,39.9764789 L6.80652562,55.6014789 L7.86602081,55.6014789 Z M18.9539599,55.6687749 L18.9539599,40.6297124 L17.9773974,40.6297124 L17.9773974,55.6687749 L18.9539599,55.6687749 Z M6.11191957,61.5486807 C10.7753172,59.0747224 11.4451975,55.1060809 11.4451975,51.6786178 L11.4451975,42.9939932 L10.3888478,42.9939932 L10.3888478,51.6786178 C10.3888478,54.8741473 9.84779058,58.404692 5.39050999,60.7755687 C5.62239164,60.9559615 5.98309643,61.3167471 6.11191957,61.5486807 Z M4.13644279,49.3033384 L4.84851961,48.2769089 C3.88212964,47.5399852 2.07650627,46.3819621 0.677783936,45.5924009 L-0.0342928868,46.4872369 C1.36442944,47.3031168 3.19548413,48.5137773 4.13644279,49.3033384 Z M1.646221,61.3164325 C2.73754339,59.0680677 4.08266168,55.8232684 5.04708612,53.2172092 L4.08266168,52.52737 C3.04209848,55.3122764 1.57008223,58.6592741 0.554898617,60.6521429 L1.646221,61.3164325 Z M14.9320882,61.4944682 L15.7696147,60.8416858 C15.008227,59.7189 13.4093128,57.9172205 12.0388149,56.6116557 L11.2774272,57.1599929 C12.6225455,58.4916691 14.1960801,60.3194598 14.9320882,61.4944682 Z M48.9701059,56.5400295 C48.7128651,56.4121357 48.2241074,56.1307693 47.9668666,55.9005605 C47.7610739,58.5095937 47.3752126,60.0443191 46.8607309,60.0187404 C44.8542523,59.9164254 43.6709443,53.7007874 43.1821867,45.1574826 L48.7900373,45.1574826 L48.7900373,43.9808598 L43.1307385,43.9808598 C43.0278422,42.2159256 42.976394,40.3998339 42.9506699,38.5070059 L41.7159138,38.5070059 C41.767362,40.3998339 41.8188102,42.2159256 41.9217065,43.9808598 L33.1497934,43.9808598 L33.1497934,45.1574826 L41.9731547,45.1574826 C42.5133605,54.5960439 43.7223925,61.4767295 46.8607309,61.5534658 C47.786798,61.5790446 48.5327965,60.4280005 48.9701059,56.5400295 Z M47.2828217,42.9961911 L48.2019394,42.2884705 C47.5891943,41.3279925 46.312642,39.7861727 45.1637449,38.6993161 L44.2956894,39.280658 C45.4190554,40.3927903 46.6956077,41.9598859 47.2828217,42.9961911 Z M31.4621015,43.3846857 L32.3709255,42.5303362 C31.5919335,41.6006031 30.0599158,40.1180555 28.7615957,39.0878107 L27.878738,39.7913925 C29.1510917,40.8718933 30.7350422,42.4046966 31.4621015,43.3846857 Z M29.1324532,60.8289063 L29.1433326,60.8166391 L29.1433326,60.8166391 L29.1655904,60.7920444 C29.1882218,60.7673796 29.2120565,60.7424444 29.2378336,60.7166978 L29.2642855,60.6906608 C29.6670719,60.2998393 30.5514511,59.6908365 34.3298709,57.0983676 C34.2012219,56.8684029 34.0211134,56.4084734 33.943924,56.1018538 L30.7791598,58.1715363 L30.7791598,46.5710938 L25.7361209,46.5710938 L25.7361209,47.7720207 L29.5955894,47.7720207 L29.5955894,57.7882617 C29.5955894,58.8614304 28.8494255,59.5002213 28.4377489,59.7557377 C28.669317,60.0368056 29.0038043,60.5478384 29.1324532,60.8289063 Z M33.9532942,59.6362488 L41.6040704,57.3589051 L41.4510549,56.2202332 L38.2377289,57.1259949 L38.2377289,50.475116 L40.9155006,50.475116 L40.9155006,49.2846863 L34.0553045,49.2846863 L34.0553045,50.475116 L37.0901124,50.475116 L37.0901124,57.4624207 C35.7639779,57.8506043 34.5398537,58.212909 33.5962579,58.4458191 L33.9532942,59.6362488 Z M66.9832549,61.3826723 L66.9832549,44.7049591 C67.5244333,43.603359 67.9883005,42.4761403 68.4263973,41.2720657 L74.1216556,41.2720657 L74.1216556,40.0936098 L60.1025583,40.0936098 L60.1025583,41.2720657 L67.0863365,41.2720657 C65.437031,45.9858894 62.7053687,50.1104852 59.2779056,52.7235831 C59.5871503,52.9797692 60.0767879,53.4152855 60.2571807,53.645853 C62.3703535,51.9037877 64.2258222,49.6237317 65.7720462,46.9850151 L65.7720462,61.3826723 L66.9832549,61.3826723 Z M51.6515551,58.6908017 C53.6477835,57.9738387 56.2326434,56.9752116 58.7151327,56.0277961 L58.4847986,54.8499283 L55.7463827,55.899767 L55.7463827,48.9093774 L58.1520939,48.9093774 L58.1520939,47.7571153 L55.7463827,47.7571153 L55.7463827,41.5092945 L58.6639473,41.5092945 L58.6639473,40.3314267 L51.3700357,40.3314267 L51.3700357,41.5092945 L54.5691197,41.5092945 L54.5691197,47.7571153 L51.7539258,47.7571153 L51.7539258,48.9093774 L54.5691197,48.9093774 L54.5691197,56.335066 C53.3406715,56.8215766 52.2145939,57.205664 51.2932577,57.5129338 L51.6515551,58.6908017 Z M72.9547403,53.5101394 L73.960209,52.747618 C72.825834,50.9938188 70.4281778,48.1724896 68.4172403,46.0882644 L67.5148965,46.7236989 C69.5000528,48.8333415 71.8461465,51.7563402 72.9547403,53.5101394 Z M98.642197,41.4653098 L98.642197,40.358539 L93.0541994,40.358539 C92.8224391,39.7150676 92.436172,38.8142076 92.0499049,38.1449973 L90.9426058,38.376647 C91.2516195,38.9943796 91.5606332,39.7150676 91.7666423,40.358539 L85.9468845,40.358539 L85.9468845,41.4653098 L98.642197,41.4653098 Z M77.2694344,57.9250172 C79.0374329,57.2335334 81.2290144,56.3779572 83.4401636,55.5121252 L84.1773838,55.2233457 L84.1773838,55.2233457 L84.6680272,55.0310292 L84.6680272,55.0310292 L84.4093351,53.9297772 L81.4861149,55.0054187 L81.4861149,45.836855 L84.1506431,45.836855 L84.1506431,44.6587714 L81.4861149,44.6587714 L81.4861149,38.5890797 L80.2702623,38.5890797 L80.2702623,44.6587714 L77.217696,44.6587714 L77.217696,45.836855 L80.2702623,45.836855 L80.2702623,55.466408 C78.976802,55.9273972 77.7868185,56.3627759 76.8555272,56.6701021 L77.2694344,57.9250172 Z M88.9756682,45.0234392 L90.1304534,44.6820526 C89.9424651,43.9467585 89.4859222,42.7912963 89.0830901,41.8984392 L87.9820159,42.1347837 C88.3579925,43.0801619 88.78768,44.2618846 88.9756682,45.0234392 Z M99.1975225,46.2773206 L99.1975225,45.1245005 L94.8843419,45.1245005 C95.2926904,44.2860859 95.7265606,43.264268 96.1349091,42.3210516 L94.9609073,41.9804456 C94.6801677,42.8712612 94.1442103,44.181284 93.7103401,45.1245005 L85.1350225,45.1245005 L85.1350225,46.2773206 L99.1975225,46.2773206 Z M84.3457585,61.8051003 C88.9302365,60.7522213 90.3210332,58.9032629 90.7331211,54.9999065 L93.3086706,54.9999065 L93.3086706,59.2114227 C93.3086706,60.4697415 93.4632035,60.8035812 93.9010469,61.034701 C94.2873794,61.2658208 94.9827777,61.3428607 95.5236431,61.3428607 L97.3265277,61.3428607 C97.8158821,61.3428607 98.5112805,61.2915008 98.8976129,61.1887809 C99.3097008,61.060381 99.5930112,60.8292612 99.7732997,60.4440616 C99.9020772,60.0331819 100.005099,58.9289429 100.030855,57.8760639 C99.6702777,57.7990239 99.2324343,57.5935841 98.9748794,57.3367844 C98.9491239,58.4410234 98.9233684,59.2884626 98.8461019,59.6736623 C98.7430799,59.981822 98.5370359,60.1872618 98.330992,60.2643017 C98.124948,60.3413416 97.6355936,60.3670216 97.2235057,60.3670216 L95.6524206,60.3670216 C95.2660881,60.3670216 94.9827777,60.3413416 94.7767337,60.2643017 C94.5449343,60.1615818 94.4934233,59.8791021 94.4934233,59.4168625 L94.4934233,54.9999065 L97.9189041,54.9999065 L97.9189041,47.9379128 L86.5349755,47.9379128 L86.5349755,54.9999065 L89.5226129,54.9999065 C89.1877914,58.3896634 87.9515277,59.930462 83.6246046,60.8292612 C83.8821596,61.034701 84.216981,61.4969406 84.3457585,61.8051003 Z M96.7542375,50.798929 L87.7698625,50.798929 L87.7698625,48.845804 L96.7542375,48.845804 L96.7542375,50.798929 Z M96.7890429,53.9613007 L87.8046679,53.9613007 L87.8046679,51.8128632 L96.7890429,51.8128632 L96.7890429,53.9613007 Z' id='测试环境' fill='%23000000' fill-rule='nonzero' opacity='0.49254441' transform='translate(49.998281, 49.975049) rotate(-30.000000) translate(-49.998281, -49.975049) '%3E%3C/path%3E%3Cg id='测试环境' opacity='0.35' fill-rule='nonzero' transform='translate(34.628798, 23.726162) rotate(-30.000000) translate(-34.628798, -23.726162) ' stroke-width='0.5'%3E%3Cuse stroke='%23575757' fill='%23000000' xlink:href='%23path-1'%3E%3C/use%3E%3Cuse stroke='%23FFFFFF' xlink:href='%23path-1'%3E%3C/use%3E%3C/g%3E%3Cg id='测试环境' opacity='0.35' fill-rule='nonzero' transform='translate(65.097548, 76.265224) rotate(-30.000000) translate(-65.097548, -76.265224) ' stroke-width='0.5'%3E%3Cuse stroke='%23575757' fill='%23000000' xlink:href='%23path-2'%3E%3C/use%3E%3Cuse stroke='%23FFFFFF' xlink:href='%23path-2'%3E%3C/use%3E%3C/g%3E%3C/g%3E%3C/g%3E%3C/svg%3E");


`

export const PILL_TAB_VALUES = {
  DBOPS: "dbops",
  RAMOPS: "ramops",
  GENERAL: "general",
  CONSOLE: "console",
  JSON_DATA: "jsonData",
  HEX_DATA: "hexData"
}

export interface Props {
  console?: string
  dbops?: DbOp[]
  ramops?: RAMOp[]
  tableops?: TableOp[]
  action: Action<any>
  traceInfo?: TraceInfo
  data: any
  displayFullContentButton: boolean
  onDisplayFullContent: () => void
  blockNum?: number
}

interface State {
  currentTab: string
  decodedDBOps: DbOp[]
  isDecodedDBOps: boolean
}

export class PillTabContentComponent extends React.Component<Props, State> {
  PILL_TABS = [{ label: t("transaction.pill.general"), value: PILL_TAB_VALUES.GENERAL }]

  constructor(props: Props) {
    super(props)

    this.state = {
      currentTab: PILL_TAB_VALUES.GENERAL,
      isDecodedDBOps: false,
      decodedDBOps: []
    }
  }

  get displayedDBOps(): DbOp[] {
    if (this.state.decodedDBOps.length > 0) {
      return this.state.decodedDBOps
    }

    if (this.props.dbops) {
      return this.props.dbops
    }

    return []
  }

  hasDBOpsToDecode() {
    return !this.state.isDecodedDBOps && this.props.dbops
  }

  onChangeContent = (currentTab: string) => {
    this.setState({ currentTab }, () => {
      if (
        this.state.currentTab === PILL_TAB_VALUES.DBOPS &&
        this.hasDBOpsToDecode() &&
        this.props.blockNum
      ) {
        decodeDBOps(this.props.dbops!, this.props.blockNum, (decodedDBOps: DbOp[]) => {
          this.setState((prevState) => ({
            currentTab: prevState.currentTab,
            decodedDBOps,
            isDecodedDBOps: true
          }))
        })
      }
    })
  }

  renderReceiverInfo() {
    if (this.props.traceInfo) {
      return (
        <DetailLine variant="compact" label={t("transaction.pill.receiver")}>
          <SearchShortcut query={`receiver:${this.props.traceInfo.receiver}`}>
            <MonospaceTextLink to={Links.viewAccount({ id: this.props.traceInfo.receiver })}>
              {this.props.traceInfo.receiver}
            </MonospaceTextLink>
          </SearchShortcut>
        </DetailLine>
      )
    }
    return null
  }

  renderAccountLink() {
    let query = `account:${this.props.action.account}`
    if (this.props.traceInfo) {
      query = `${query} receiver:${this.props.traceInfo.receiver}`
    }
    return (
      <DetailLine variant="compact" label={t("transaction.pill.account")}>
        <SearchShortcut query={query}>
          <MonospaceTextLink to={Links.viewAccount({ id: this.props.action.account })}>
            {this.props.action.account}
          </MonospaceTextLink>{" "}
        </SearchShortcut>
      </DetailLine>
    )
  }

  renderActionName() {
    let query = `action:${this.props.action.name} account:${this.props.action.account}`
    if (this.props.traceInfo) {
      query = `${query} receiver:${this.props.traceInfo.receiver}`
    }
    return (
      <DetailLine variant="compact" label={t("transaction.pill.action_name")}>
        <SearchShortcut query={query}>
          <Text>{this.props.action.name}</Text>
        </SearchShortcut>
      </DetailLine>
    )
  }

  renderAuthorizations() {
    const authorizations = (this.props.action.authorization || []).map(
      (entry: Authorization, index: number) => {
        return (
          <Cell key={index}>
            <SearchShortcut query={`auth:${entry.actor}@${entry.permission}`}>
              <AutorizationBox authorization={entry} />
            </SearchShortcut>
          </Cell>
        )
      }
    )

    return (
      <DetailLine variant="compact" label={t("transaction.pill.authorization")}>
        <Text>{authorizations}</Text>
      </DetailLine>
    )
  }

  renderDisplayFullContentButton() {
    return this.props.displayFullContentButton ? (
      <Cell float="right">
        <LinkStyledText color={theme.colors.link} onClick={() => this.props.onDisplayFullContent()}>
          Show Full Content
        </LinkStyledText>
      </Cell>
    ) : null
  }

  renderTabContent() {
    if (this.state.currentTab === PILL_TAB_VALUES.GENERAL) {
      return (
        <ContentWrapper>
          {this.renderReceiverInfo()}
          {this.renderAccountLink()}
          {this.renderActionName()}
          {this.renderAuthorizations()}
        </ContentWrapper>
      )
    }

    if (this.state.currentTab === PILL_TAB_VALUES.JSON_DATA) {
      return (
        <RawWrapper px="24px" py="10px">
          {this.renderDisplayFullContentButton()}
          <JsonWrapper>{JSON.stringify(this.props.data, null, "   ")}</JsonWrapper>
        </RawWrapper>
      )
    }

    if (this.state.currentTab === PILL_TAB_VALUES.HEX_DATA) {
      return (
        <RawWrapper px="24px" py="10px">
          <JsonWrapper>{this.props.action.hex_data}</JsonWrapper>
        </RawWrapper>
      )
    }

    if (this.state.currentTab === PILL_TAB_VALUES.RAMOPS) {
      return (
        <ContentWrapper>
          <RamUsage type="detailed" ramops={this.props.ramops || []} />
        </ContentWrapper>
      )
    }

    if (this.state.currentTab === PILL_TAB_VALUES.DBOPS) {
      return (
        <ContentWrapper>
          <DBOperations tableops={this.props.tableops || []} dbops={this.displayedDBOps || []} />
        </ContentWrapper>
      )
    }

    if (this.state.currentTab === PILL_TAB_VALUES.CONSOLE) {
      return (
        <ContentWrapper>
          <JsonWrapper>{this.props.console!.replace(/\\r/g, "")}</JsonWrapper>
        </ContentWrapper>
      )
    }

    return null
  }

  render() {
    const tabs = [...this.PILL_TABS]

    if (this.props.action.data) {
      tabs.push({ label: t("transaction.pill.jsonData"), value: PILL_TAB_VALUES.JSON_DATA })
    } else if (this.props.action.hex_data) {
      tabs.push({ label: t("transaction.pill.hexData"), value: PILL_TAB_VALUES.HEX_DATA })
    }

    if (this.props.dbops && this.props.dbops.length > 0) {
      tabs.push({ label: t("transaction.pill.dbOps"), value: PILL_TAB_VALUES.DBOPS })
    }

    if (this.props.ramops && this.props.ramops.length > 0) {
      tabs.push({ label: t("transaction.pill.ramOps"), value: PILL_TAB_VALUES.RAMOPS })
    }

    if (this.props.console && this.props.console.length > 0) {
      tabs.push({ label: t("transaction.pill.console"), value: PILL_TAB_VALUES.CONSOLE })
    }

    return [
      <VerticalTabs key={1} tabData={tabs} onSelectTab={this.onChangeContent} />,
      <TabContentWrapper key={2}>{this.renderTabContent()}</TabContentWrapper>
    ]
  }
}
