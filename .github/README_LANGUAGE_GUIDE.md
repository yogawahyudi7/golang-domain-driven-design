# README Language Management

This document provides guidance for maintaining multiple language versions of the README.

## 📋 Available Languages

- �� **English** (Primary): `README.md`
- �� **Indonesian**: `README.id.md`

## 🔄 Sync Guidelines

When updating README content:

1. **Primary Language (English)**: Update `README.md` first
2. **Secondary Languages**: Update corresponding language files
3. **Keep Structure Consistent**: Maintain same section order and anchor links
4. **Update Navigation**: Ensure language toggle links work properly

## 📝 Translation Checklist

### Content Sections
- [ ] Project description
- [ ] Table of contents
- [ ] Quick start guide
- [ ] Installation instructions
- [ ] Configuration details
- [ ] API documentation references
- [ ] Contributing guidelines
- [ ] License information

### UI Elements
- [ ] Language toggle buttons
- [ ] Navigation links
- [ ] Call-to-action buttons
- [ ] Footer links
- [ ] Badge descriptions

### Technical Content
- [ ] Code examples (keep in English)
- [ ] Command line instructions
- [ ] Configuration files
- [ ] Error messages
- [ ] Log outputs

## 🌐 Adding New Languages

To add a new language (e.g., Japanese):

1. Create `README.ja.md`
2. Translate content from `README.md`
3. Update language toggle in all README files:
   ```markdown
   [�� English](README.md) • [�� Bahasa Indonesia](README.id.md) • [🇯🇵 日本語](README.ja.md)
   ```
4. Update this template file

## 🎯 Best Practices

- **Use Consistent Emojis**: Keep the same emoji patterns across languages
- **Maintain Anchor Links**: Ensure internal links work in all versions
- **Code Stays English**: Keep code examples, commands, and technical terms in English
- **Cultural Adaptation**: Adapt examples and references to be culturally relevant
- **Version Sync**: Always sync all language versions when making structural changes

## 🛠️ Automation Ideas

Future improvements could include:
- GitHub Actions to detect README changes
- Automated translation validation
- Consistency checking between language versions
- Translation status tracking
