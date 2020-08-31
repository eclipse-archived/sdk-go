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
	return CreateSelector([]string{fo.prefix, "fdu", fduid, "records", "*", "info"})
}

// GetAllFDUsRecordsSelector ...
func (fo *FOrchestrator) GetAllFDUsRecordsSelector() *yaks.Selector {
	return CreateSelector([]string{fo.prefix, "fdu", "*", "records", "*", "info"})
}

// ID Extraction

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
func (fo *FOrchestrator) GetFDUInfo(fduid string) (*FOrCEFDUDescriptor, error) {
	s, _ := yaks.NewSelector(fo.GetFDUPath(fduid).ToString())
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return nil, &FError{"FDU not found", nil}
	}
	v := kvs[0].Value().ToString()
	sv := FOrCEFDUDescriptor{}
	err := json.Unmarshal([]byte(v), &sv)
	if err != nil {
		return nil, err
	}
	return &sv, nil
}

// GetAllFDUsInfo ...
func (fo *FOrchestrator) GetAllFDUsInfo() ([]FOrCEFDUDescriptor, error) {
	s := fo.GetAllFDUsSelector()
	kvs := fo.ws.Get(s)
	if len(kvs) == 0 {
		return []FOrCEFDUDescriptor{}, &FError{"No entities found", nil}
	}
	var entities []FOrCEFDUDescriptor = []FOrCEFDUDescriptor{}
	for _, kv := range kvs {
		v := kv.Value().ToString()
		sv := FOrCEFDUDescriptor{}
		err := json.Unmarshal([]byte(v), &sv)
		if err != nil {
			return entities, err
		}
		entities = append(entities, sv)
	}
	return entities, nil
}

// AddFDUInfo ...
func (fo *FOrchestrator) AddFDUInfo(info FOrCEFDUDescriptor) error {
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
func (fo *FOrchestrator) ObserveFDUs(listener func(FOrCEFDUDescriptor)) (*yaks.SubscriptionID, error) {
	s := fo.GetAllFDUsSelector()

	cb := func(kvs []yaks.Change) {
		if len(kvs) > 0 {
			v := kvs[0].Value().ToString()
			sv := FOrCEFDUDescriptor{}
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
