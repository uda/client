// @flow
import * as React from 'react'
import * as Types from '../../constants/types/fs'
import * as Kb from '../../common-adapters'
import * as Styles from '../../styles'
import * as Flow from '../../util/flow'
import {isMobile} from '../../constants/platform'

type Props = {
  path: Types.Path,
  reason: Types.SoftError,
}

const Explain = (props: Props) => {
  const elems = Types.getPathElements(props.path)
  if (elems.length < 3) {
    return null
  }
  switch (elems[1]) {
    case 'public':
      return null
    case 'private':
      return (
        <Kb.Box2 direction="horizontal" style={styles.explainBox}>
          <Kb.Text center={true} type="Body">
            Only people in the private folder can access this.
          </Kb.Text>
        </Kb.Box2>
      )
    case 'team':
      return (
        <Kb.Box2 direction="horizontal" style={styles.explainBox}>
          <Kb.Text center={true} type="Body">
            Only members of
          </Kb.Text>
          <Kb.Text type="BodySemibold" style={styles.explainTextTeam}>
            {elems[2]}
          </Kb.Text>
          <Kb.Text center={true} type="Body">
            can access this.
          </Kb.Text>
        </Kb.Box2>
      )
    default:
      return null
  }
}

const NoAccess = (props: Props) => (
  <Kb.Box2 direction="vertical" style={styles.container} fullWidth={true}>
    <Kb.Box2 direction="vertical" style={styles.main} fullWidth={true} centerChildren={true}>
      <Kb.Icon
        type={isMobile ? 'icon-fancy-no-access-mobile-128-125' : 'icon-fancy-no-access-desktop-96-94'}
      />
      <Kb.Text type="Header" style={styles.textYouDontHave}>
        You don't have access to this folder or file.
      </Kb.Text>
      <Explain {...props} />
    </Kb.Box2>
  </Kb.Box2>
)

const NonExistent = (props: Props) => (
  <Kb.Box2 direction="vertical" style={styles.container} fullWidth={true}>
    <Kb.Box2 direction="vertical" style={styles.main} fullWidth={true} centerChildren={true}>
      <Kb.Icon
        type={
          isMobile
            ? 'icon-fancy-folder-file-inexistant-mobile-188-120'
            : 'icon-fancy-folder-file-inexistant-desktop-153-94'
        }
      />
      <Kb.Text type="Header" style={styles.textYouDontHave}>
        This file or folder doesn't exist.
      </Kb.Text>
      <Kb.Box2 direction="horizontal" style={styles.explainBox}>
        <Kb.Text center={true} type="Body">
          Either it was deleted, or the path is incorrect.
        </Kb.Text>
      </Kb.Box2>
    </Kb.Box2>
  </Kb.Box2>
)

const Oops = (props: Props) => {
  switch (props.reason) {
    case 'no-access':
      return <NoAccess {...props} />
    case 'non-existent':
      return <NonExistent {...props} />
    default:
      Flow.ifFlowComplainsAboutThisFunctionYouHaventHandledAllCasesInASwitch(props.reason)
      return null
  }
}

export default Oops

const styles = Styles.styleSheetCreate({
  container: Styles.platformStyles({
    common: {
      ...Styles.globalStyles.flexGrow,
      backgroundColor: Styles.globalColors.white,
    },
    isMobile: {
      padding: Styles.globalMargins.large,
    },
  }),
  explainBox: Styles.platformStyles({
    isElectron: {
      marginTop: Styles.globalMargins.small,
    },
    isMobile: {
      marginTop: Styles.globalMargins.medium,
    },
  }),
  explainTextTeam: {
    marginLeft: Styles.globalMargins.xtiny,
    marginRight: Styles.globalMargins.xtiny,
  },
  footer: {
    paddingBottom: Styles.globalMargins.large,
  },
  header: {
    backgroundColor: Styles.globalColors.red,
    height: 40,
  },
  main: {
    ...Styles.globalStyles.flexGrow,
  },
  textYouDontHave: Styles.platformStyles({
    isElectron: {
      marginTop: Styles.globalMargins.medium,
    },
    isMobile: {
      marginTop: Styles.globalMargins.xlarge,
      textAlign: 'center',
    },
  }),
})
