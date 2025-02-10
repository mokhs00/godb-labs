// Code generated by BobGen mysql v0.30.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

type Factory struct {
	basePostMods PostModSlice
	baseUserMods UserModSlice
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) NewPost(mods ...PostMod) *PostTemplate {
	o := &PostTemplate{f: f}

	if f != nil {
		f.basePostMods.Apply(o)
	}

	PostModSlice(mods).Apply(o)

	return o
}

func (f *Factory) NewUser(mods ...UserMod) *UserTemplate {
	o := &UserTemplate{f: f}

	if f != nil {
		f.baseUserMods.Apply(o)
	}

	UserModSlice(mods).Apply(o)

	return o
}

func (f *Factory) ClearBasePostMods() {
	f.basePostMods = nil
}

func (f *Factory) AddBasePostMod(mods ...PostMod) {
	f.basePostMods = append(f.basePostMods, mods...)
}

func (f *Factory) ClearBaseUserMods() {
	f.baseUserMods = nil
}

func (f *Factory) AddBaseUserMod(mods ...UserMod) {
	f.baseUserMods = append(f.baseUserMods, mods...)
}
