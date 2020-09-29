/*
* Copyright (c) 2014,2019 Contributors to the Eclipse Foundation
* See the NOTICE file(s) distributed with this work for additional
* information regarding copyright ownership.
* This program and the accompanying materials are made available under the
* terms of the Eclipse Public License 2.0 which is available at
* http://www.eclipse.org/legal/epl-2.0, or the Apache License, Version 2.0
* which is available at https://www.apache.org/licenses/LICENSE-2.0.
* SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
* Contributors: Gabriele Baldoni, ADLINK Technology Inc.
* golang APIs
 */

package fog05sdk

import (
	"encoding/json"
	"strings"

	"github.com/atolab/yaks-go"
)

const (
	// FIM is FIM EntityKind
	//FIM string = "FIM"

	//OrchestratorPrefix is the prefix for the orchestrator in Zenoh
	OrchestratorPrefix string = "/force"
)

// Job represents a job inside the orchestrator
type Job struct {
	JobID          string `json:"job_id"`
	OriginalSender string `json:"original_sender"`
	JobKind        string `json:"kind"`
	Body           string `json:"body"`
	Status         string `json:"status"`
}

// RequestNewJobMessage represents a request to do a new job for the orchestrator
type RequestNewJobMessage struct {
	Sender  string `json:"sender"`
	JobKind string `json:"job_kind"`
	Body    string `json:"body"`
}

// ReplyNewJobMessage represents a reply to do a new job request from the orchestrator
type ReplyNewJobMessage struct {
	OriginalSender string  `json:"original_sender"`
	Accepted       bool    `json:"accepted"`
	JobID          string  `json:"job_id"`
	Body           *string `json:"body,omitempty"`
}

//FIMInfo represents information required for connection with a FIM
type FIMInfo struct {
	UUID    string `json:"uuid"`
	Locator string `json:"locator"`
}

//CloudInfo represents information required for connection with a Cloud (K8s)
type CloudInfo struct {
	UUID   string `json:"uuid"`
	Config string `json:"config"`
}

// FOrchestrator ...
type FOrchestrator struct {
	ws        *yaks.Workspace
	prefix    string
	listeners []*yaks.SubscriptionID
	evals     []*yaks.Path
}

// NewFOrchestrator ...
func NewFOrchestrator(wspace *yaks.Workspace, prefix string) FOrchestrator {
	return FOrchestrator{ws: wspace, prefix: prefix, listeners: []*yaks.SubscriptionID{}, evals: []*yaks.Path{}}
}

// Unsubscribe ...
func (fo *FOrchestrator) Unsubscribe(sid *yaks.SubscriptionID) error {
	err := fo.ws.Unsubscribe(sid)
	if err != nil {
		return err
	}
	p := -1
	for i, e := range fo.listeners {
		if e == sid {
			p = i
		}
	}
	if p == -1 {
		return &FError{"Subscriber not found!!", nil}
	}
	fo.listeners = append(fo.listeners[:p], fo.listeners[p+1:]...)
	return nil

}

// RemoveEval ...
func (fo *FOrchestrator) RemoveEval(sid *yaks.Path) error {
	err := fo.ws.UnregisterEval(sid)
	if err != nil {
		return err
	}
	p := -1

	for i, e := range fo.evals {
		if e == sid {
			p = i
		}
	}
	if p == -1 {
		return &FError{"Eval not found!!", nil}
	}
	fo.evals = append(fo.evals[:p], fo.evals[p+1:]...)
	return nil

}

// Path+Selector Generation

// GetJobRequestSelector ...
func (fo *FOrchestrator) GetJobRequestSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "newjob"})
}

// GetJobInfoPath ...
func (fo *FOrchestrator) GetJobInfoPath(jobid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "job", jobid, "info"})
}

// GetAllJobsInfoSelector ...
func (fo *FOrchestrator) GetAllJobsInfoSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "job", "*", "info"})
}

// Catalog Paths and Selectors

// GetFIMPath ...
func (fo *FOrchestrator) GetFIMPath(fimid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "fim", fimid, "info"})
}

// GetAllFIMsSelector ...
func (fo *FOrchestrator) GetAllFIMsSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "fim", "*", "info"})
}

// GetCloudPath ...
func (fo *FOrchestrator) GetCloudPath(cloudid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "cloud", cloudid, "info"})
}

// GetAllCloudsSelector ...
func (fo *FOrchestrator) GetAllCloudsSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "cloud", "*", "info"})
}

// GetEntityPath ...
func (fo *FOrchestrator) GetEntityPath(entityid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "entity", entityid, "info"})
}

// GetAllEntitiesSelector ...
func (fo *FOrchestrator) GetAllEntitiesSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "entity", "*", "info"})
}

// GetEntityRecordPath ...
func (fo *FOrchestrator) GetEntityRecordPath(entityid string, instanceid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "entity", entityid, "record", instanceid, "info"})
}

// GetEntityRecordSelector ...
func (fo *FOrchestrator) GetEntityRecordSelector(instanceid string) *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "entity", "*", "record", instanceid, "info"})
}

// GetAllEntityRecordsSelector ...
func (fo *FOrchestrator) GetAllEntityRecordsSelector(entityid string) *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "entity", entityid, "record", "*", "info"})
}

// GetVirtualLinkPath ...
func (fo *FOrchestrator) GetVirtualLinkPath(vlid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "vl", vlid, "info"})
}

// GetAllVirtualLinksSelector ...
func (fo *FOrchestrator) GetAllVirtualLinksSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "vl", "*", "info"})
}

// GetVirtualLinkRecordPath ...
func (fo *FOrchestrator) GetVirtualLinkRecordPath(vlid string, instanceid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "vl", vlid, "record", instanceid, "info"})
}

// GetAllVirtualLinkRecordsSelector ...
func (fo *FOrchestrator) GetAllVirtualLinkRecordsSelector(vlid string) *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "vl", vlid, "record", "*", "info"})
}

// GetFDUPath ...
func (fo *FOrchestrator) GetFDUPath(fduid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "fdu", fduid, "info"})
}

// GetAllFDUsSelector ...
func (fo *FOrchestrator) GetAllFDUsSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "fdu", "*", "info"})
}

// GetFDURecordPath ...
func (fo *FOrchestrator) GetFDURecordPath(fduid string, instanceid string) *yaks.Path {
	return CreatePath([]string{fo.prefix, "fdu", fduid, "record", instanceid, "info"})
}

// GetAllFDURecordsSelector ...
func (fo *FOrchestrator) GetAllFDURecordsSelector(fduid string) *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "fdu", fduid, "record", "*", "info"})
}

// GetAllFDUsRecordsSelector ...
func (fo *FOrchestrator) GetAllFDUsRecordsSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "fdu", "*", "record", "*", "info"})
}

// ID Extraction

// ExtractFIMID ...
func (fo *FOrchestrator) ExtractFIMID(path *yaks.Path) string {
	return strings.Split(path.ToString(), URISeparator)[4]
}

// ExtractCloudID ...
func (fo *FOrchestrator) ExtractCloudID(path *yaks.Path) string {
	return strings.Split(path.ToString(), URISeparator)[4]
}

// ExtractEntityID ...
func (fo *FOrchestrator) ExtractEntityID(path *yaks.Path) string {
	return strings.Split(path.ToString(), URISeparator)[4]
}

// ExtractEntityInstanceID ...
func (fo *FOrchestrator) ExtractEntityInstanceID(path *yaks.Path) string {
	return strings.Split(path.ToString(), URISeparator)[4]
}

// ExtractNetworkID ...
func (fo *FOrchestrator) ExtractNetworkID(path *yaks.Path) string {
	return strings.Split(path.ToString(), URISeparator)[4]
}

// ExtractNetworkInstanceID ...
func (fo *FOrchestrator) ExtractNetworkInstanceID(path *yaks.Path) string {
	return strings.Split(path.ToString(), URISeparator)[4]
}

// ExtractFDUID ...
func (fo *FOrchestrator) ExtractFDUID(path *yaks.Path) string {
	return strings.Split(path.ToString(), URISeparator)[4]
}

// ExtractFDUInstanceID ...
func (fo *FOrchestrator) ExtractFDUInstanceID(path *yaks.Path) string {
	return strings.Split(path.ToString(), URISeparator)[4]
}

// Get/Put/Subscribe/Evals

// AddJobInfo ...
func (fo *FOrchestrator) AddJobInfo(info Job) error {
	s := fo.GetJobInfoPath(info.JobID)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	sv := yaks.NewStringValue(string(v))
	err = fo.ws.Put(s, sv)
	return err
}

//GetJobInfo ...
func (fo *FOrchestrator) GetJobInfo(jobid string) (*Job, error) {
	s, _ := yaks.NewSelector(fo.GetJobInfoPath(jobid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"Entity not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := Job{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllJobsInfo ...
func (fo *FOrchestrator) GetAllJobsInfo() ([]Job, error) {
	s := fo.GetAllJobsInfoSelector()
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []Job{}, &FError{"No jobs found", nil}
	}
	var jobs []Job = []Job{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := Job{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return jobs, err
		}
		jobs = append(jobs, sv)
	}
	return jobs, nil
}

//RemoveJobInfo ...
func (fo *FOrchestrator) RemoveJobInfo(jobid string) error {
	s := fo.GetJobInfoPath(jobid)
	err := fo.ws.Remove(s)
	return err
}

// GetFIMInfo ...
func (fo *FOrchestrator) GetFIMInfo(fimid string) (*FIMInfo, error) {
	s, _ := yaks.NewSelector(fo.GetFIMPath(fimid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"Entity not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := FIMInfo{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllFIMsInfo ...
func (fo *FOrchestrator) GetAllFIMsInfo() ([]FIMInfo, error) {
	s := fo.GetAllFIMsSelector()
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []FIMInfo{}, &FError{"No fims found", nil}
	}
	var fims []FIMInfo = []FIMInfo{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := FIMInfo{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return fims, err
		}
		fims = append(fims, sv)
	}
	return fims, nil
}

// AddFIMInfo ...
func (fo *FOrchestrator) AddFIMInfo(info FIMInfo) error {
	s := fo.GetFIMPath(info.UUID)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	sv := yaks.NewStringValue(string(v))
	err = fo.ws.Put(s, sv)
	return err
}

// RemoveFIMInfo ...
func (fo *FOrchestrator) RemoveFIMInfo(fimid string) error {
	s := fo.GetFIMPath(fimid)
	err := fo.ws.Remove(s)
	return err
}

// ObserveFIMs ...
func (fo *FOrchestrator) ObserveFIMs(listener func(FIMInfo)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllFIMsSelector()

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := FIMInfo{}
			err := json.Unmarshal([]byte(v), &sv)
			if err != nil {
				panic(err.Error())
			}
			listener(sv)
		}
	}

	sid, err := fo.ws.Subscribe(s, cb)
	if err != nil {
		return nil, err
	}
	fo.listeners = append(fo.listeners, sid)
	return sid, nil
}

// GetCloudInfo ...
func (fo *FOrchestrator) GetCloudInfo(cloudid string) (*CloudInfo, error) {
	s, _ := yaks.NewSelector(fo.GetCloudPath(cloudid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"Entity not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := CloudInfo{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllCloudsInfo ...
func (fo *FOrchestrator) GetAllCloudsInfo() ([]CloudInfo, error) {
	s := fo.GetAllCloudsSelector()
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []CloudInfo{}, &FError{"No clouds found", nil}
	}
	var clouds []CloudInfo = []CloudInfo{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := CloudInfo{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return clouds, err
		}
		clouds = append(clouds, sv)
	}
	return clouds, nil
}

// AddCloudInfo ...
func (fo *FOrchestrator) AddCloudInfo(info CloudInfo) error {
	s := fo.GetCloudPath(info.UUID)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	sv := yaks.NewStringValue(string(v))
	err = fo.ws.Put(s, sv)
	return err
}

// RemoveCloudInfo ...
func (fo *FOrchestrator) RemoveCloudInfo(cloudid string) error {
	s := fo.GetCloudPath(cloudid)
	err := fo.ws.Remove(s)
	return err
}

// ObserveClouds ...
func (fo *FOrchestrator) ObserveClouds(listener func(CloudInfo)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllCloudsSelector()

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := CloudInfo{}
			err := json.Unmarshal([]byte(v), &sv)
			if err != nil {
				panic(err.Error())
			}
			listener(sv)
		}
	}

	sid, err := fo.ws.Subscribe(s, cb)
	if err != nil {
		return nil, err
	}
	fo.listeners = append(fo.listeners, sid)
	return sid, nil
}

// GetEntityInfo ...
func (fo *FOrchestrator) GetEntityInfo(fduid string) (*EntityDescriptor, error) {
	s, _ := yaks.NewSelector(fo.GetEntityPath(fduid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"Entity not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := EntityDescriptor{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllEntitiesInfo ...
func (fo *FOrchestrator) GetAllEntitiesInfo() ([]EntityDescriptor, error) {
	s, _ := yaks.NewSelector(fo.GetAllEntitiesSelector().ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []EntityDescriptor{}, &FError{"No entities found", nil}
	}
	var entities []EntityDescriptor = []EntityDescriptor{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := EntityDescriptor{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return entities, err
		}
		entities = append(entities, sv)
	}
	return entities, nil
}

// AddEntityInfo ...
func (fo *FOrchestrator) AddEntityInfo(info EntityDescriptor) error {
	s := fo.GetEntityPath(*info.UUID)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	sv := yaks.NewStringValue(string(v))
	err = fo.ws.Put(s, sv)
	return err
}

// RemoveEntityInfo ...
func (fo *FOrchestrator) RemoveEntityInfo(entityid string) error {
	s := fo.GetEntityPath(entityid)
	err := fo.ws.Remove(s)
	return err
}

// ObserveEntities ...
func (fo *FOrchestrator) ObserveEntities(listener func(EntityDescriptor)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllEntitiesSelector()

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := EntityDescriptor{}
			err := json.Unmarshal([]byte(v), &sv)
			if err != nil {
				panic(err.Error())
			}
			listener(sv)
		}
	}

	sid, err := fo.ws.Subscribe(s, cb)
	if err != nil {
		return nil, err
	}
	fo.listeners = append(fo.listeners, sid)
	return sid, nil
}

// GetEntityInstanceInfo ...
func (fo *FOrchestrator) GetEntityInstanceInfo(entityid string, instanceid string) (*EntityRecord, error) {
	s, _ := yaks.NewSelector(fo.GetEntityRecordPath(entityid, instanceid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"Entity Instance not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := EntityRecord{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// FindEntityInstanceInfo ...
func (fo *FOrchestrator) FindEntityInstanceInfo(instanceid string) (*EntityRecord, error) {
	s := fo.GetEntityRecordSelector(instanceid)
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"Entity Instance not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := EntityRecord{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllEntityRecordsInfo ...
func (fo *FOrchestrator) GetAllEntityRecordsInfo(entityid string) ([]EntityRecord, error) {
	s := fo.GetAllEntityRecordsSelector(entityid)
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []EntityRecord{}, &FError{"No entities found", nil}
	}
	var entities []EntityRecord = []EntityRecord{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := EntityRecord{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return entities, err
		}
		entities = append(entities, sv)
	}
	return entities, nil
}

// AddEntityRecord ...
func (fo *FOrchestrator) AddEntityRecord(info EntityRecord) error {
	s := fo.GetEntityRecordPath(info.ID, info.UUID)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	sv := yaks.NewStringValue(string(v))
	err = fo.ws.Put(s, sv)
	return err
}

// RemoveEntityRecord ...
func (fo *FOrchestrator) RemoveEntityRecord(entityid string, instanceid string) error {
	s := fo.GetEntityRecordPath(entityid, instanceid)
	err := fo.ws.Remove(s)
	return err
}

// ObserveEntityRecords ...
func (fo *FOrchestrator) ObserveEntityRecords(entityid string, listener func(EntityRecord)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllEntityRecordsSelector(entityid)

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := EntityRecord{}
			err := json.Unmarshal([]byte(v), &sv)
			if err != nil {
				panic(err.Error())
			}
			listener(sv)
		}
	}

	sid, err := fo.ws.Subscribe(s, cb)
	if err != nil {
		return nil, err
	}
	fo.listeners = append(fo.listeners, sid)
	return sid, nil
}

// GetVirtualLinkInfo ...
func (fo *FOrchestrator) GetVirtualLinkInfo(vlid string) (*VirtualLinkDescriptor, error) {
	s, _ := yaks.NewSelector(fo.GetVirtualLinkPath(vlid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"Entity not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := VirtualLinkDescriptor{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllVirtualLinksInfo ...
func (fo *FOrchestrator) GetAllVirtualLinksInfo() ([]VirtualLinkDescriptor, error) {
	s := fo.GetAllVirtualLinksSelector()
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []VirtualLinkDescriptor{}, &FError{"No entities found", nil}
	}
	var entities []VirtualLinkDescriptor = []VirtualLinkDescriptor{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := VirtualLinkDescriptor{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return entities, err
		}
		entities = append(entities, sv)
	}
	return entities, nil
}

// AddVirtualLinkInfo ...
func (fo *FOrchestrator) AddVirtualLinkInfo(info VirtualLinkDescriptor) error {
	s := fo.GetVirtualLinkPath(*info.UUID)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	sv := yaks.NewStringValue(string(v))
	err = fo.ws.Put(s, sv)
	return err
}

// RemoveVirtualLinkInfo ...
func (fo *FOrchestrator) RemoveVirtualLinkInfo(vlid string) error {
	s := fo.GetVirtualLinkPath(vlid)
	err := fo.ws.Remove(s)
	return err
}

// ObserveVirtualLinks ...
func (fo *FOrchestrator) ObserveVirtualLinks(listener func(VirtualLinkDescriptor)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllVirtualLinksSelector()

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := VirtualLinkDescriptor{}
			err := json.Unmarshal([]byte(v), &sv)
			if err != nil {
				panic(err.Error())
			}
			listener(sv)
		}
	}

	sid, err := fo.ws.Subscribe(s, cb)
	if err != nil {
		return nil, err
	}
	fo.listeners = append(fo.listeners, sid)
	return sid, nil
}

// GetFDUInfo ...
func (fo *FOrchestrator) GetFDUInfo(fduid string) (*FOrcEFDUDescriptor, error) {
	s, _ := yaks.NewSelector(fo.GetFDUPath(fduid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"FDU not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := FOrcEFDUDescriptor{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllFDUsInfo ...
func (fo *FOrchestrator) GetAllFDUsInfo() ([]FOrcEFDUDescriptor, error) {
	s := fo.GetAllFDUsSelector()
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []FOrcEFDUDescriptor{}, &FError{"No entities found", nil}
	}
	var entities []FOrcEFDUDescriptor = []FOrcEFDUDescriptor{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := FOrcEFDUDescriptor{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return entities, err
		}
		entities = append(entities, sv)
	}
	return entities, nil
}

// AddFDUInfo ...
func (fo *FOrchestrator) AddFDUInfo(info FOrcEFDUDescriptor) error {
	s := fo.GetFDUPath(*info.UUID)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	sv := yaks.NewStringValue(string(v))
	err = fo.ws.Put(s, sv)
	return err
}

// RemoveFDUInfo ...
func (fo *FOrchestrator) RemoveFDUInfo(fduid string) error {
	s := fo.GetFDUPath(fduid)
	err := fo.ws.Remove(s)
	return err
}

// ObserveFDUs ...
func (fo *FOrchestrator) ObserveFDUs(listener func(FOrcEFDUDescriptor)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllFDUsSelector()

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := FOrcEFDUDescriptor{}
			err := json.Unmarshal([]byte(v), &sv)
			if err != nil {
				panic(err.Error())
			}
			listener(sv)
		}
	}

	sid, err := fo.ws.Subscribe(s, cb)
	if err != nil {
		return nil, err
	}
	fo.listeners = append(fo.listeners, sid)
	return sid, nil
}

// GetFDURecord ...
func (fo *FOrchestrator) GetFDURecord(fduid string, instanceid string) (*FOrcEFDURecord, error) {
	s, _ := yaks.NewSelector(fo.GetFDURecordPath(fduid, instanceid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"FDU not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := FOrcEFDURecord{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllFDURecordsInfo ...
func (fo *FOrchestrator) GetAllFDURecordsInfo(fduid string) ([]FOrcEFDURecord, error) {
	s := fo.GetAllFDURecordsSelector(fduid)
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []FOrcEFDURecord{}, &FError{"No entities found", nil}
	}
	var entities []FOrcEFDURecord = []FOrcEFDURecord{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := FOrcEFDURecord{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return entities, err
		}
		entities = append(entities, sv)
	}
	return entities, nil
}

// AddFDURecord ...
func (fo *FOrchestrator) AddFDURecord(info FOrcEFDURecord) error {
	s := fo.GetFDURecordPath(info.ID, info.UUID)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	sv := yaks.NewStringValue(string(v))
	err = fo.ws.Put(s, sv)
	return err
}

// RemoveFDURecord ...
func (fo *FOrchestrator) RemoveFDURecord(fduid string, instanceid string) error {
	s := fo.GetFDURecordPath(fduid, instanceid)
	err := fo.ws.Remove(s)
	return err
}

// ObserveFDURecords ...
func (fo *FOrchestrator) ObserveFDURecords(fduid string, listener func(FOrcEFDURecord)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllFDURecordsSelector(fduid)

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := FOrcEFDURecord{}
			err := json.Unmarshal([]byte(v), &sv)
			if err != nil {
				panic(err.Error())
			}
			listener(sv)
		}
	}

	sid, err := fo.ws.Subscribe(s, cb)
	if err != nil {
		return nil, err
	}
	fo.listeners = append(fo.listeners, sid)
	return sid, nil
}

// ObserveFDUsRecords ...
func (fo *FOrchestrator) ObserveFDUsRecords(listener func(FOrcEFDURecord)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllFDUsRecordsSelector()

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := FOrcEFDURecord{}
			err := json.Unmarshal([]byte(v), &sv)
			if err != nil {
				panic(err.Error())
			}
			listener(sv)
		}
	}

	sid, err := fo.ws.Subscribe(s, cb)
	if err != nil {
		return nil, err
	}
	fo.listeners = append(fo.listeners, sid)
	return sid, nil
}

// FOrcEZConnector ...
type FOrcEZConnector struct {
	zclient      *yaks.Yaks
	zadmin       *yaks.Admin
	ws           *yaks.Workspace
	Orchestrator FOrchestrator
}

// NewFOrcEZConnector ...
func NewFOrcEZConnector(locator string, sysid string, tenantid string) (*FOrcEZConnector, error) {
	z, err := yaks.Login(&locator, nil)
	if err != nil {
		return nil, err
	}

	ad := z.Admin()

	wpath, err := yaks.NewPath("/")
	if err != nil {
		return nil, err
	}

	ws := z.WorkspaceWithExecutor(wpath)

	p := OrchestratorPrefix + URISeparator + sysid +
		URISeparator + "tenant" + URISeparator + tenantid

	o := NewFOrchestrator(ws, p)

	return &FOrcEZConnector{ws: ws, Orchestrator: o, zadmin: ad, zclient: z}, nil
}

// Close closes the connector
func (fzc *FOrcEZConnector) Close() error {
	return fzc.zclient.Logout()
}
