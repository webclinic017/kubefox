{{- define "labels" -}}
{{ include "selectors" . }}
app.kubernetes.io/managed-by: {{ printf "%s-operator" .Instance.Name | cleanLabel | quote }}
kubefox.xigxog.io/runtime-version: {{ .BuildInfo.Version | cleanLabel | quote }}
{{- if .Component.IsPlatformComponent }}
  {{- with .Component.Hash }}
kubefox.xigxog.io/component-hash: {{ . | cleanLabel | quote }}
  {{- end }}
{{- end }}
{{- with .Component.Type }}
kubefox.xigxog.io/component-type: {{ . | cleanLabel | quote }}
{{- end }}
{{- with .Component.Hash }}
kubefox.xigxog.io/component-hash-short: {{ . | substr 0 7 | cleanLabel | quote }}
{{- end }}
{{- range $k, $v := .Component.Labels }}
{{ $k }}: {{ $v | cleanLabel | quote }}
{{- end }}
{{- end }}

{{- define "annotations" -}}
{{- with .Hash }}
kubefox.xigxog.io/template-data-hash: {{ . | quote }}
{{- end }}
{{- with .Component.Annotations }}
{{ . | toYaml }}
{{- end }}
{{- end }}

{{- define "selectors" -}}
app.kubernetes.io/instance: {{ .Instance.Name | cleanLabel | quote }}
{{- with .Platform.Name }}
kubefox.xigxog.io/platform: {{ . | cleanLabel | quote }}
{{- end }}
{{- with .Component.Name }}
app.kubernetes.io/component: {{ . | cleanLabel | quote }}
{{- end }}
{{- if not .Component.IsPlatformComponent }}
  {{- with .Component.Hash }}
kubefox.xigxog.io/component-hash: {{ . | cleanLabel | quote }}
  {{- end }}
{{- end }}
{{- end }}

{{- define "metadata" -}}
metadata:
  name: {{ name }}
  namespace: {{ .Platform.Namespace }}
  labels:
    {{- include "labels" . | nindent 4 }}
  annotations:
    {{- include "annotations" . | nindent 4 }}
  {{- with .Owner }}
  ownerReferences:
    {{- . | toYaml | nindent 4 }}
  {{- end }}
{{- end }}

{{- define "roleBinding" -}}
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
{{ include "metadata" . }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: {{ name }}
subjects:
  - kind: ServiceAccount
    name: {{ name }}
    namespace: {{ .Platform.Namespace }}
{{- end }}

{{- define "clusterRoleBinding" -}}
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
{{ include "metadata" . }}
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: {{ name }}
subjects:
  - kind: ServiceAccount
    name: {{ name }}
    namespace: {{ .Platform.Namespace }}
{{- end }}

{{- define "serviceAccount" -}}
apiVersion: v1
kind: ServiceAccount
{{ include "metadata" . }}
{{- end }}

{{- define "env" -}}
{{- with .Component.Name }}
- name: KUBEFOX_COMPONENT
  value: {{ . | quote }}
{{- end }}
{{- with .Component.Hash }}
- name: KUBEFOX_HASH
  value: {{ . | quote }}
{{- end }}
{{- with .Component.Type }}
- name: KUBEFOX_COMPONENT_TYPE
  value: {{ . | quote }}
{{- end }}
- name: KUBEFOX_HOST_IP
  valueFrom:
    fieldRef:
      fieldPath: status.hostIP
- name: KUBEFOX_NODE
  valueFrom:
    fieldRef:
      fieldPath: spec.nodeName
- name: KUBEFOX_POD
  valueFrom:
    fieldRef:
      fieldPath: metadata.name
- name: KUBEFOX_POD_IP
  valueFrom:
    fieldRef:
      fieldPath: status.podIP
{{- with .Values.GOMEMLIMIT }}
- name: GOMEMLIMIT
  value: {{ . | quote }}
{{- end }}
{{- with .Values.GOMAXPROCS }}
- name: GOMAXPROCS
  value: {{ . | quote }}
{{- end }}
{{- end }}

{{- define "podSpec" -}}
serviceAccountName: {{ name }}
securityContext:
  runAsNonRoot: true
  runAsUser: 100
  runAsGroup: 1000
  fsGroup: 1000
  fsGroupChangePolicy: OnRootMismatch

{{- with .Component.ImagePullSecret }}
imagePullSecrets:
  - name: {{ . }}
{{- end }}

{{- with .Component.NodeSelector }}
nodeSelector:
  {{- . | toYaml | nindent 2 }}
{{- end }}

{{- with .Component.NodeName }}
nodeName: {{ . | quote  }}
{{- end }}

{{- with .Component.Affinity }}
affinity:
  {{- . | toYaml | nindent 2 }}
{{- end }}

{{- with .Component.Tolerations }}
tolerations:
  {{- . | toYaml | nindent 2 }}
{{- end }}
{{- end }}

{{- define "securityContext" -}}
securityContext:
  allowPrivilegeEscalation: false
  capabilities:
    drop:
      - ALL
{{- end }}

{{- define "resources" -}}
{{- with .Component.Resources }}
resources:
  {{- . | toYaml | nindent 2 }}
{{- end }}
{{- end }}

{{- define "probes" -}}
{{- with .Component.LivenessProbe }}
livenessProbe:
  {{- . | toYaml | nindent 2 }}
{{- end }}
{{- with .Component.ReadinessProbe }}
readinessProbe:
  {{- . | toYaml | nindent 2 }}
{{- end }}
{{- with .Component.StartupProbe }}
startupProbe:
  {{- . | toYaml | nindent 2 }}
{{- end }}
{{- end }}

{{- define "bootstrap" -}}
name: bootstrap
image: {{ .Instance.BootstrapImage }}
imagePullPolicy: {{ .Component.ImagePullPolicy | default "IfNotPresent" }}
{{ include "securityContext" . }}
args:
  - -instance={{ .Instance.Name }}
  - -platform-namespace={{ .Platform.Namespace }}
  - -component={{ .Component.Name }}
  - -component-service-name={{ printf "%s.%s" name .Platform.Namespace }}
  - -component-ip=$(KUBEFOX_COMPONENT_IP)
  - -vault-url={{ .Values.vaultURL }}
  - -log-format={{ .Telemetry.Logs.Format | default "json" }}
  - -log-level={{ .Telemetry.Logs.Level | default "info" }}
env:
{{- include "env" . | nindent 2 }}
  - name: KUBEFOX_COMPONENT_IP
    valueFrom:
      fieldRef:
        fieldPath: status.podIP
envFrom:
  - configMapRef:
      name: {{ .Platform.Name }}-env
volumeMounts:
  - name: root-ca
    mountPath: {{ homePath }}/ca.crt
    subPath: ca.crt
  - name: kubefox
    mountPath: {{ homePath }}
{{- end }}